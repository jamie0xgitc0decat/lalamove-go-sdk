package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
)

// QuotationService handles quotation-related operations
type QuotationService struct {
	Client *Client
}

// Create creates a new quotation
func (s *QuotationService) Create(ctx context.Context, req *models.QuotationRequest) (*models.QuotationData, error) {
	// Validate request
	if err := s.validateQuotationRequest(req); err != nil {
		return nil, err
	}

	var response models.QuotationResponse
	err := s.client.Request().Do(ctx, http.MethodPost, "/quotations", req, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create quotation: %w", err)
	}

	return &response.Data, nil
}

// Get retrieves a quotation's details by ID
func (s *QuotationService) Get(ctx context.Context, quotationID string) (*models.QuotationData, error) {
	if quotationID == "" {
		return nil, fmt.Errorf("quotation ID is required")
	}

	path := fmt.Sprintf("/quotations/%s", quotationID)
	var response models.QuotationResponse
	err := s.client.Request().Do(ctx, http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get quotation details: %w", err)
	}

	return &response.Data, nil
}

func (s *QuotationService) validateQuotationRequest(req *models.QuotationRequest) error {
	if req == nil {
		return fmt.Errorf("quotation request cannot be nil")
	}

	if req.ServiceType == "" {
		return fmt.Errorf("service type is required")
	}

	if req.Language == "" {
		return fmt.Errorf("language is required")
	}

	if len(req.Stops) < 2 {
		return fmt.Errorf("minimum 2 stops required")
	}

	if len(req.Stops) > 16 {
		return fmt.Errorf("maximum 16 stops allowed")
	}

	// Validate each stop
	for i, stop := range req.Stops {
		if stop.Address == "" {
			return fmt.Errorf("address is required for stop %d", i+1)
		}

		if stop.Coordinates.Latitude == 0 || stop.Coordinates.Longitude == 0 {
			return fmt.Errorf("coordinates are required for stop %d", i+1)
		}
	}

	return nil
}
