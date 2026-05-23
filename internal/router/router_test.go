package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/handler"
	"github.com/SiddhantAgarwal/GoServerTemplate/internal/service"
)

type mockHealthService struct{}

func (m *mockHealthService) Check(ctx context.Context) (map[string]any, error) {
	return map[string]any{"status": "ok"}, nil
}

func TestIndexRoute(t *testing.T) {
	hh := handler.NewHealthHandler(service.HealthService(&mockHealthService{}))
	r := NewRouter(hh)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}

	if rr.Body.String() == "" {
		t.Error("expected non-empty body")
	}
}

func TestHealthRoute(t *testing.T) {
	hh := handler.NewHealthHandler(service.HealthService(&mockHealthService{}))
	r := NewRouter(hh)

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}
}

func TestNotFound(t *testing.T) {
	hh := handler.NewHealthHandler(service.HealthService(&mockHealthService{}))
	r := NewRouter(hh)

	req, _ := http.NewRequest(http.MethodGet, "/unknown", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404 got %d", rr.Code)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	hh := handler.NewHealthHandler(service.HealthService(&mockHealthService{}))
	r := NewRouter(hh)

	req, _ := http.NewRequest(http.MethodPost, "/health", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405 got %d", rr.Code)
	}
}
