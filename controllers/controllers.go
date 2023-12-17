package controllers
// Requried Imports
import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
	"time"
)

type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

// Control Function For Single Process
func ProcessSingle(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Failed to decode JSON payload. Ensure the payload is valid JSON.", http.StatusBadRequest)
		return
	}

	startTime := time.Now()
	sortedArrays := make([][]int, len(payload.ToSort))
	for i, arr := range payload.ToSort {
		sortedArrays[i] = make([]int, len(arr))
		copy(sortedArrays[i], arr)
		sort.Ints(sortedArrays[i])
	}
	timeTaken := time.Since(startTime)

	response := map[string]interface{}{
		"sorted_arrays": sortedArrays,
		"time_ns":       float64(timeTaken.Nanoseconds()) / float64(time.Second), // Convert to seconds
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response. Contact the API administrator.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Control Function For Concurrent Process
func ProcessConcurrent(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Failed to decode JSON payload. Ensure the payload is valid JSON.", http.StatusBadRequest)
		return
	}

	startTime := time.Now()
	var wg sync.WaitGroup
	sortedArrays := make([][]int, len(payload.ToSort))

	for i, arr := range payload.ToSort {
		wg.Add(1)
		go func(i int, arr []int) {
			defer wg.Done()
			sortedArrays[i] = make([]int, len(arr))
			copy(sortedArrays[i], arr)
			sort.Ints(sortedArrays[i])
		}(i, arr)
	}

	wg.Wait()
	timeTaken := time.Since(startTime)

	response := map[string]interface{}{
		"sorted_arrays": sortedArrays,
		"time_ns":       float64(timeTaken.Nanoseconds()) / float64(time.Second), // Convert to seconds
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response. Contact the API administrator.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
