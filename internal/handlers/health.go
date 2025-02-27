package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse representa a resposta do health check
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler retorna o status da API
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
