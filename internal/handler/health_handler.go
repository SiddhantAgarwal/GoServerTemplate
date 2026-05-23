package handler

import (
	"net/http"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/service"
)

type HealthHandler struct {
	svc service.HealthService
}

func NewHealthHandler(svc service.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

func (h *HealthHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	resp, err := h.svc.Check(r.Context())
	if err != nil {
		JSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	JSON(w, http.StatusOK, resp)
}

func (h *HealthHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, map[string]bool{"isServerDown": false})
}
