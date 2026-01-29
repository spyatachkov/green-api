package v1

import (
	"net/http"

	"github.com/spyatachkov/green-api/backend/internal/api/dto"
	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (h *Handler) GetStateInstance(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")

	if idInstance == "" || apiToken == "" {
		h.sendError(w, "idInstance and apiToken are required", http.StatusBadRequest)
		return
	}

	creds := &models.GreenAPICredentials{
		IDInstance:       idInstance,
		APITokenInstance: apiToken,
	}

	serviceResp, err := h.service.GetStateInstance(r.Context(), creds)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.FromService(serviceResp)
	h.sendJSON(w, response, http.StatusOK)
}
