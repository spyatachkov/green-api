package dto

import (
	"github.com/spyatachkov/green-api/backend/internal/models"
	"github.com/spyatachkov/green-api/backend/pkg/whatsapp"
)

func ToSendMessageRequest(m *models.SendMessageRequest) *SendMessageRequest {
	return &SendMessageRequest{
		ChatID:  whatsapp.FormatChatID(m.ChatID),
		Message: m.Message,
	}
}

func ToSendFileRequest(m *models.SendFileRequest) *SendFileRequest {
	return &SendFileRequest{
		ChatID:   whatsapp.FormatChatID(m.ChatID),
		URLFile:  m.FileURL,
		FileName: m.FileName,
	}
}

func ToServiceResponse(data interface{}, err error, statusCode int) *models.GreenAPIResponse {
	success := err == nil && statusCode >= 200 && statusCode < 300

	resp := &models.GreenAPIResponse{
		Success: success,
		Data:    data,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	return resp
}
