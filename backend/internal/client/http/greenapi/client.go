package greenapi

import (
	"net/http"
	"time"

	httpclient "github.com/spyatachkov/green-api/backend/internal/client/http"
)

type greenApiClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string, timeout time.Duration) httpclient.GreenApiClient {
	return &greenApiClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
