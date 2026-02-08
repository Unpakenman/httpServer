package clinics

import (
	"encoding/json"
	"httpServer/internal/app/httpserver/models"
	"net/http"
)

func (r *httpRouter) Ping(w http.ResponseWriter, req *http.Request) {
	response := models.PingResponse{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	r.logger.Info("Получен входящий запрос /ping")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
