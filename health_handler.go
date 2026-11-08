package main

import (
	"net/http"

	"github.com/Elanoej/service-code-api/utils"
)

type HealthHandler struct {
}

type healthResponse struct {
	Status string `json:"status"`
}

func buildHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Registrar(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.health)
}

func (h *HealthHandler) health(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, http.StatusOK, healthResponse{Status: "ok"})
}
