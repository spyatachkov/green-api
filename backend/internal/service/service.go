package service

import (
	"context"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

type GreenAPIService interface {
	GetSettings(ctx context.Context, creds *models.GreenAPICredentials) (*models.GreenAPIResponse, error)
	GetStateInstance(ctx context.Context, creds *models.GreenAPICredentials) (*models.GreenAPIResponse, error)
	SendMessage(ctx context.Context, req *models.SendMessageRequest) (*models.GreenAPIResponse, error)
	SendFileByUrl(ctx context.Context, req *models.SendFileRequest) (*models.GreenAPIResponse, error)
}
