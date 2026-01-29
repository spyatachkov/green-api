package v1

import (
	"encoding/json"
	"net/http"

	"github.com/spyatachkov/green-api/backend/internal/api/dto"
	"github.com/spyatachkov/green-api/backend/internal/service"
)

type Handler struct {
	service service.GreenAPIService
}

func NewHandler(svc service.GreenAPIService) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) sendError(w http.ResponseWriter, message string, statusCode int) {
	resp := dto.ErrorResponse{Error: message}
	h.sendJSON(w, resp, statusCode)
}
