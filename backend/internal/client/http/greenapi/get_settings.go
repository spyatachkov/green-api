package greenapi

import (
	"context"
	"fmt"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (c *greenApiClient) GetSettings(ctx context.Context, creds *models.GreenAPICredentials) (*models.GreenAPIResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getSettings/%s", c.baseURL, creds.IDInstance, creds.APITokenInstance)
	return c.makeRequest(ctx, "GET", url, nil)
}
