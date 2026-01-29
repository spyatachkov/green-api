package greenapi

import (
	"context"
	"fmt"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (c *greenApiClient) GetStateInstance(ctx context.Context, creds *models.GreenAPICredentials) (*models.GreenAPIResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", c.baseURL, creds.IDInstance, creds.APITokenInstance)
	return c.makeRequest(ctx, "GET", url, nil)
}
