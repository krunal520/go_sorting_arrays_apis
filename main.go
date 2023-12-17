// main.go
package main

import (
	"fmt"
	"mygoapp/routes"
	"net/http"
)

func main() {
	router := routes.InitRoutes()

	port := 8000
	serverAddr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server is listening on %s...\n", serverAddr)
	http.ListenAndServe(serverAddr, router)
}
