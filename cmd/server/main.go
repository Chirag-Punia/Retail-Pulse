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
