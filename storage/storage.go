package storage

import (
	"sync"
	"testtask/models"
	"time"
)

type RequestStorage struct {
	mu      sync.Mutex
	request []models.Request
}

func NewRequest() *RequestStorage {
	return &RequestStorage{
		request: make([]models.Request, 0),
	}
}

func (s *RequestStorage) AddRequest(req models.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.request = append(s.request, req)
}

func (s *RequestStorage) GetRequestsSince(timeLimit time.Time) []models.Request {
	s.mu.Lock()
	defer s.mu.Unlock()

	var filteredRequests []models.Request
	for _, req := range s.request {
		if req.Time.After(timeLimit) {
			filteredRequests = append(filteredRequests, req)
		}
	}

	return filteredRequests
}
