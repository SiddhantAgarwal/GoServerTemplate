package service

import (
	"context"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/repository"
)

type HealthService interface {
	Check(ctx context.Context) (map[string]any, error)
}

type healthService struct {
	repo repository.HealthRepository
}

func NewHealthService(repo repository.HealthRepository) HealthService {
	return &healthService{repo: repo}
}

func (s *healthService) Check(ctx context.Context) (map[string]any, error) {
	if err := s.repo.Ping(ctx); err != nil {
		return map[string]any{"status": "unhealthy", "db": err.Error()}, nil
	}

	return map[string]any{"status": "ok"}, nil
}
