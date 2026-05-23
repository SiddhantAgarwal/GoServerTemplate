package service

import (
	"context"
	"errors"
	"testing"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/repository"
)

type mockRepo struct {
	err error
}

func (m *mockRepo) Ping(ctx context.Context) error {
	return m.err
}

func TestHealthServiceOK(t *testing.T) {
	repo := &mockRepo{err: nil}
	svc := NewHealthService(repository.HealthRepository(repo))

	resp, err := svc.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp["status"] != "ok" {
		t.Errorf("expected status ok, got %v", resp["status"])
	}
}

func TestHealthServiceUnhealthy(t *testing.T) {
	repo := &mockRepo{err: errors.New("connection refused")}
	svc := NewHealthService(repository.HealthRepository(repo))

	resp, err := svc.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp["status"] != "unhealthy" {
		t.Errorf("expected status unhealthy, got %v", resp["status"])
	}
}
