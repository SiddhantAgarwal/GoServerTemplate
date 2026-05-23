package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/service"
)

type mockSvc struct {
	resp map[string]any
	err  error
}

func (m *mockSvc) Check(ctx context.Context) (map[string]any, error) {
	return m.resp, m.err
}

func TestHealthHandlerOK(t *testing.T) {
	svc := &mockSvc{resp: map[string]any{"status": "ok"}}
	hh := NewHealthHandler(service.HealthService(svc))

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()
	hh.HandleHealth(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}

	if !contains(rr.Body.String(), "ok") {
		t.Errorf("expected body to contain 'ok', got %s", rr.Body.String())
	}
}

func TestHealthHandlerError(t *testing.T) {
	svc := &mockSvc{err: errors.New("boom")}
	hh := NewHealthHandler(service.HealthService(svc))

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()
	hh.HandleHealth(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 got %d", rr.Code)
	}
}

func TestIndexHandler(t *testing.T) {
	hh := NewHealthHandler(service.HealthService(&mockSvc{}))

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	hh.HandleIndex(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}

	if !contains(rr.Body.String(), "isServerDown") {
		t.Errorf("expected body to contain 'isServerDown', got %s", rr.Body.String())
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && findSubstr(s, substr))
}

func findSubstr(s, substr string) bool {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}

	return false
}
