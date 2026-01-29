package v1

import (
	"encoding/json"
	"net/http"

	"github.com/spyatachkov/green-api/backend/internal/api/dto"
	"github.com/spyatachkov/green-api/backend/pkg/validator"
)

func (h *Handler) SendFileByUrl(w http.ResponseWriter, r *http.Request) {
	var apiReq dto.SendFileRequest
	if err := json.NewDecoder(r.Body).Decode(&apiReq); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if apiReq.IDInstance == "" || apiReq.APITokenInstance == "" {
		h.sendError(w, "idInstance and apiTokenInstance are required", http.StatusBadRequest)
		return
	}

	if apiReq.ChatID == "" || apiReq.FileURL == "" {
		h.sendError(w, "chatId and fileUrl are required", http.StatusBadRequest)
		return
	}

	if err := validator.ValidatePhone(apiReq.ChatID); err != nil {
		h.sendError(w, "Invalid phone number: "+err.Error(), http.StatusBadRequest)
		return
	}

	serviceReq := apiReq.ToSendFileService()

	serviceResp, err := h.service.SendFileByUrl(r.Context(), serviceReq)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.FromService(serviceResp)
	h.sendJSON(w, response, http.StatusOK)
}
