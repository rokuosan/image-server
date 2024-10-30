package handler

import "net/http"

type HealthHandler struct{}

func (h *HealthHandler) Get(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, `{"status": "ok"}`)
}
