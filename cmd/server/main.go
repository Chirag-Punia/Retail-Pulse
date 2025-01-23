package main

import (
	"log"
	"net/http"

	"retail-pulse/internal/handlers"
	"retail-pulse/internal/processor"
	"retail-pulse/internal/store"

	"github.com/gorilla/mux"
)

func main() {

	jobStore := store.NewJobStore()
	imageProcessor := processor.NewImageProcessor()
	handler := handlers.NewHandler(jobStore, imageProcessor)

	r := mux.NewRouter()

	r.HandleFunc("/api/submit", handler.SubmitJob).Methods("POST")
	r.HandleFunc("/api/status", handler.GetJobStatus).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}
