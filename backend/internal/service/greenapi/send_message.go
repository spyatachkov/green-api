package greenapi

import (
	"context"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (s *svc) SendMessage(ctx context.Context, req *models.SendMessageRequest) (*models.GreenAPIResponse, error) {
	return s.client.SendMessage(ctx, req)
}
