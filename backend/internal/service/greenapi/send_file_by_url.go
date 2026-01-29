package greenapi

import (
	"context"
	"fmt"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (s *svc) SendFileByUrl(ctx context.Context, req *models.SendFileRequest) (*models.GreenAPIResponse, error) {
	phone := req.ChatID
	for _, c := range phone {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("phone must contain only digits")
		}
	}

	return s.client.SendFileByUrl(ctx, &models.SendFileRequest{
		Credentials: req.Credentials,
		ChatID:      phone,
		FileURL:     req.FileURL,
		FileName:    req.FileName,
	})
}
