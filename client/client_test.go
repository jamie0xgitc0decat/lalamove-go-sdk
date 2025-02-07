package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		expected *Client
	}{
		{
			name: "Default configuration",
			config: Config{
				APIKey:    "test-key",
				APISecret: "test-secret",
			},
			expected: &Client{
				apiKey:      "test-key",
				apiSecret:   "test-secret",
				environment: Sandbox,
				baseURL:     Sandbox.GetBaseURL(),
			},
		},
		{
			name: "Production configuration",
			config: Config{
				APIKey:      "test-key",
				APISecret:   "test-secret",
				Environment: Production,
			},
			expected: &Client{
				apiKey:      "test-key",
				apiSecret:   "test-secret",
				environment: Production,
				baseURL:     Production.GetBaseURL(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.config)
			assert.NotNil(t, client)
			assert.Equal(t, tt.expected.apiKey, client.apiKey)
			assert.Equal(t, tt.expected.apiSecret, client.apiSecret)
			assert.Equal(t, tt.expected.environment, client.environment)
			assert.Equal(t, tt.expected.baseURL, client.baseURL)
		})
	}
}

func TestClient_IsProduction(t *testing.T) {
	sandboxClient := NewClient(Config{
		Environment: Sandbox,
	})
	assert.False(t, sandboxClient.IsProduction())

	prodClient := NewClient(Config{
		Environment: Production,
	})
	assert.True(t, prodClient.IsProduction())
}

func TestClientOptions(t *testing.T) {
	customTimeout := 60 * time.Second
	customBaseURL := "https://custom.api.url"

	client := NewClient(
		"test-key",
		"test-secret",
		WithTimeout(customTimeout),
		WithBaseURL(customBaseURL),
	)

	assert.Equal(t, customTimeout, client.httpClient.Timeout)
	assert.Equal(t, customBaseURL, client.baseURL)
}

func TestOrderService_Create(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v3/orders", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)

		// Mock response
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{
			"id": "test-order-id",
			"status": "ASSIGNING_DRIVER",
			"created_at": "2024-03-14T12:00:00Z"
		}`))
	}))
	defer server.Close()

	// Create client with test server URL
	client := NewClient(
		"test-key",
		"test-secret",
		WithBaseURL(server.URL),
	)

	order, err := client.Orders.Create(context.Background(), &models.Order{
		ServiceType: "MOTORCYCLE",
		Stops: []models.Stop{
			{
				Address: "Test Address",
				Coordinates: models.Coordinates{
					Latitude:  1.234,
					Longitude: 4.567,
				},
			},
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "test-order-id", order.ID)
}
