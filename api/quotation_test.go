package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestQuotationService_Get(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request
		assert.Equal(t, "/v3/quotations/test-quotation-id", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		// Return mock response
		scheduleAt := time.Date(2022, 4, 13, 7, 18, 38, 0, time.UTC)
		expiresAt := time.Date(2022, 4, 13, 7, 23, 39, 0, time.UTC)

		response := models.QuotationResponse{
			Data: models.QuotationData{
				QuotationID:     "test-quotation-id",
				ScheduleAt:      scheduleAt,
				ExpiresAt:       expiresAt,
				ServiceType:     "MOTORCYCLE",
				SpecialRequests: []string{"TOLL_FEE_10"},
				Language:        "EN_HK",
				Stops: []models.QuotationStop{
					{
						StopID: "stop-1",
						Coordinates: models.Coordinates{
							Latitude:  22.3354735,
							Longitude: 114.1761581,
						},
						Address: "Test Address 1",
					},
					{
						StopID: "stop-2",
						Coordinates: models.Coordinates{
							Latitude:  22.2812946,
							Longitude: 114.1598610,
						},
						Address: "Test Address 2",
					},
				},
				PriceBreakdown: models.PriceBreakdown{
					Base:            "90",
					SpecialRequests: "13",
					VAT:             "21",
					Total:           "124",
					Currency:        "HKD",
				},
			},
		}

		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with test server
	client := NewClient(Config{
		APIKey:      "test-key",
		APISecret:   "test-secret",
		Environment: "sandbox",
		Market:      "HK",
		BaseURL:     server.URL,
	})

	// Test Get quotation
	quotation, err := client.Quotations.Get(context.Background(), "test-quotation-id")
	assert.NoError(t, err)
	assert.NotNil(t, quotation)
	assert.Equal(t, "test-quotation-id", quotation.QuotationID)
	assert.Equal(t, "MOTORCYCLE", quotation.ServiceType)
	assert.Equal(t, "124", quotation.PriceBreakdown.Total)
	assert.Equal(t, "HKD", quotation.PriceBreakdown.Currency)
}

func TestQuotationService_Get_ValidationError(t *testing.T) {
	client := NewClient(Config{
		APIKey:    "test-key",
		APISecret: "test-secret",
		Market:    "HK",
	})

	quotation, err := client.Quotations.Get(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, quotation)
	assert.Contains(t, err.Error(), "quotation ID is required")
}
