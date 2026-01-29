package greenapi

import (
	"context"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (s *svc) GetStateInstance(ctx context.Context, creds *models.GreenAPICredentials) (*models.GreenAPIResponse, error) {
	return s.client.GetStateInstance(ctx, creds)
}
