package greenapi

import (
	httpclient "github.com/spyatachkov/green-api/backend/internal/client/http"
	"github.com/spyatachkov/green-api/backend/internal/service"
)

type svc struct {
	client httpclient.GreenApiClient
}

func New(client httpclient.GreenApiClient) service.GreenAPIService {
	return &svc{
		client: client,
	}
}
