package main

import (
	"net/http"

	v1 "github.com/ethirajmudhaliar/backend/react-go-gov-search-data/govData/v1"
	"github.com/ethirajmudhaliar/backend/react-go-gov-search-data/logger"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// LoggingMiddleware logs details about incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Define the routes for the API
	router.HandleFunc("/api/data", v1.GetGovernmentData).Methods("GET")

	// Add the logging middleware
	router.Use(LoggingMiddleware)

	return router
}

func main() {
	router := SetupRouter()

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://react-frontend", "http://localhost:3000"}, // Frontend URL
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow all headers
	}).Handler(router)

	logger.Info("Starting server on port 8080")

	err := http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		logger.Error("Error starting server: " + err.Error())
	}
}
