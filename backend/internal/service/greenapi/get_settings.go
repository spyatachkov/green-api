package greenapi

import (
	"context"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (s *svc) GetSettings(ctx context.Context, creds *models.GreenAPICredentials) (*models.GreenAPIResponse, error) {
	return s.client.GetSettings(ctx, creds)
}
