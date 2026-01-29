package greenapi

import (
	"context"
	"fmt"

	"github.com/spyatachkov/green-api/backend/internal/client/http/greenapi/dto"
	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (c *greenApiClient) SendFileByUrl(ctx context.Context, req *models.SendFileRequest) (*models.GreenAPIResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s",
		c.baseURL,
		req.Credentials.IDInstance,
		req.Credentials.APITokenInstance,
	)

	clientDTO := dto.ToSendFileRequest(req)

	return c.makeRequest(ctx, "POST", url, clientDTO)
}
