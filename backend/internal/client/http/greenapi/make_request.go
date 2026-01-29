package greenapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

func (c *greenApiClient) makeRequest(ctx context.Context, method, url string, payload interface{}) (*models.GreenAPIResponse, error) {
	var body io.Reader

	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return &models.GreenAPIResponse{
				Success: false,
				Error:   fmt.Sprintf("Error marshaling request: %v", err),
			}, err
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return &models.GreenAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Error creating request: %v", err),
		}, err
	}

	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &models.GreenAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Error making request: %v", err),
		}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &models.GreenAPIResponse{
			Success: false,
			Error:   fmt.Sprintf("Error reading response: %v", err),
		}, err
	}

	var data interface{}
	if err := json.Unmarshal(respBody, &data); err != nil {
		data = string(respBody)
	}

	success := resp.StatusCode >= 200 && resp.StatusCode < 300

	return &models.GreenAPIResponse{
		Success: success,
		Data:    data,
		Error: func() string {
			if !success {
				return fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody))
			}
			return ""
		}(),
	}, nil
}
