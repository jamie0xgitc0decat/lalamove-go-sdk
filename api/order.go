package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/interfaces"
	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
)

type OrderService struct {
	Client interfaces.Client
}

// Create creates a new order
func (s *OrderService) Create(ctx context.Context, order *models.Order) (*models.Order, error) {
	var response models.Order
	err := s.Client.Request().Do(ctx, http.MethodPost, "/orders", order, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Get retrieves an order by ID
func (s *OrderService) Get(ctx context.Context, orderID string) (*models.Order, error) {
	var response models.Order
	path := fmt.Sprintf("/orders/%s", orderID)
	err := s.Client.Request().Do(ctx, http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
