package v1

import (
	"encoding/json"
	"net/http"

	"github.com/spyatachkov/green-api/backend/internal/api/dto"
	"github.com/spyatachkov/green-api/backend/pkg/validator"
)

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var apiReq dto.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&apiReq); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if apiReq.IDInstance == "" || apiReq.APITokenInstance == "" {
		h.sendError(w, "idInstance and apiTokenInstance are required", http.StatusBadRequest)
		return
	}

	if apiReq.ChatID == "" || apiReq.Message == "" {
		h.sendError(w, "chatId and message are required", http.StatusBadRequest)
		return
	}

	if err := validator.ValidatePhone(apiReq.ChatID); err != nil {
		h.sendError(w, "Invalid phone number: "+err.Error(), http.StatusBadRequest)
		return
	}

	serviceReq := apiReq.ToSendMessageService()

	serviceResp, err := h.service.SendMessage(r.Context(), serviceReq)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.FromService(serviceResp)
	h.sendJSON(w, response, http.StatusOK)
}
