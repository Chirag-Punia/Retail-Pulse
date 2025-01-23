package handlers

import (
	"encoding/json"
	"net/http"

	"retail-pulse/internal/models"
	"retail-pulse/internal/processor"
	"retail-pulse/internal/store"

	"github.com/google/uuid"
)

type Handler struct {
	jobStore       *store.JobStore
	imageProcessor *processor.ImageProcessor
}

func NewHandler(jobStore *store.JobStore, imageProcessor *processor.ImageProcessor) *Handler {
	return &Handler{
		jobStore:       jobStore,
		imageProcessor: imageProcessor,
	}
}

func (h *Handler) SubmitJob(w http.ResponseWriter, r *http.Request) {
	var req models.JobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Count != len(req.Visits) {
		http.Error(w, "Count does not match number of visits", http.StatusBadRequest)
		return
	}

	jobID := uuid.New().String()

	h.jobStore.Create(jobID, models.JobStatus{
		Status: models.StatusOngoing,
		JobID:  jobID,
	})

	go h.processJob(jobID, req)

	response := map[string]string{
		"job_id": jobID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetJobStatus(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, "Job ID is required", http.StatusBadRequest)
		return
	}

	status, exists := h.jobStore.Get(jobID)
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct{}{})
		return
	}

	json.NewEncoder(w).Encode(status)
}

func (h *Handler) processJob(jobID string, req models.JobRequest) {
	var errors []models.JobError

	for _, visit := range req.Visits {
		for _, imageURL := range visit.ImageURLs {
			result, err := h.imageProcessor.ProcessImage(imageURL)
			if err != nil {
				errors = append(errors, models.JobError{
					StoreID: visit.StoreID,
					Error:   err.Error(),
				})
				break
			}

			_ = result
		}
	}

	status := models.StatusCompleted
	if len(errors) > 0 {
		status = models.StatusFailed
	}

	h.jobStore.Update(jobID, models.JobStatus{
		Status: status,
		JobID:  jobID,
		Errors: errors,
	})
}
