package dto

import "github.com/spyatachkov/green-api/backend/internal/models"

func (r *SendMessageRequest) ToSendMessageService() *models.SendMessageRequest {
	return &models.SendMessageRequest{
		Credentials: models.GreenAPICredentials{
			IDInstance:       r.IDInstance,
			APITokenInstance: r.APITokenInstance,
		},
		ChatID:  r.ChatID,
		Message: r.Message,
	}
}

func (r *SendFileRequest) ToSendFileService() *models.SendFileRequest {
	return &models.SendFileRequest{
		Credentials: models.GreenAPICredentials{
			IDInstance:       r.IDInstance,
			APITokenInstance: r.APITokenInstance,
		},
		ChatID:   r.ChatID,
		FileURL:  r.FileURL,
		FileName: r.FileName,
	}
}

func FromService(serviceResp *models.GreenAPIResponse) *Response {
	return &Response{
		Success: serviceResp.Success,
		Data:    serviceResp.Data,
		Error:   serviceResp.Error,
	}
}
