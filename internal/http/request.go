package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/google/uuid"
	"github.com/jamie0xgitc0decat/lalamove-go-sdk/internal/auth"
)

// Request handles HTTP requests to the Lalamove API
type Request struct {
	client      *http.Client
	credentials auth.Credentials
	baseURL     string
	market      string
}

// NewRequest creates a new request handler
func NewRequest(client *http.Client, credentials auth.Credentials, baseURL, market string) *Request {
	return &Request{
		client:      client,
		credentials: credentials,
		baseURL:     baseURL,
		market:      market,
	}
}

// Do executes an HTTP request and decodes the response
func (r *Request) Do(ctx context.Context, method, endpoint string, body, result interface{}) error {
	// Create request URL
	url := path.Join(r.baseURL, endpoint)

	// Marshal body if present
	var bodyBytes []byte
	var err error
	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Sign request
	token, _, err := auth.SignRequest(r.credentials, method, endpoint, bodyBytes)
	if err != nil {
		return fmt.Errorf("failed to sign request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("hmac %s", token))
	req.Header.Set("Market", r.market)
	req.Header.Set("Request-ID", uuid.New().String())
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code
	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Decode response if result interface is provided
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
