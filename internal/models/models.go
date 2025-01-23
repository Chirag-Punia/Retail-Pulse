package models

const (
	StatusOngoing   = "ongoing"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type JobRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

type JobResponse struct {
	JobID string `json:"job_id"`
}

type JobStatus struct {
	Status  string        `json:"status"`
	JobID   string        `json:"job_id"`
	Errors  []JobError    `json:"errors,omitempty"`
	Results []ImageResult `json:"results,omitempty"`
}

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

type ImageResult struct {
	StoreID string `json:"store_id"`
	URL     string `json:"url"`
	Detail  string `json:"detail"`
}
