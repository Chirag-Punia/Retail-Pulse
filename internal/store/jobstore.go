package store

import (
	"sync"

	"retail-pulse/internal/models"
)

type JobStore struct {
	sync.RWMutex
	jobs map[string]models.JobStatus
}

func NewJobStore() *JobStore {
	return &JobStore{
		jobs: make(map[string]models.JobStatus),
	}
}

func (s *JobStore) Create(jobID string, status models.JobStatus) {
	s.Lock()
	defer s.Unlock()
	s.jobs[jobID] = status
}

func (s *JobStore) Update(jobID string, status models.JobStatus) {
	s.Lock()
	defer s.Unlock()
	s.jobs[jobID] = status
}

func (s *JobStore) Get(jobID string) (models.JobStatus, bool) {
	s.RLock()
	defer s.RUnlock()
	status, exists := s.jobs[jobID]
	return status, exists
}
