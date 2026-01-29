package greenapi

import (
	"context"
	"fmt"

	"github.com/spyatachkov/green-api/backend/internal/client/http/greenapi/dto"
	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (c *greenApiClient) SendMessage(ctx context.Context, req *models.SendMessageRequest) (*models.GreenAPIResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s",
		c.baseURL,
		req.Credentials.IDInstance,
		req.Credentials.APITokenInstance,
	)

	clientDTO := dto.ToSendMessageRequest(req)

	return c.makeRequest(ctx, "POST", url, clientDTO)
}
