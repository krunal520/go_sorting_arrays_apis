package routes

import (
	"mygoapp/controllers"

	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes for the application
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/process-single", controllers.ProcessSingle).Methods("POST")
	router.HandleFunc("/process-concurrent", controllers.ProcessConcurrent).Methods("POST")

	return router
}
