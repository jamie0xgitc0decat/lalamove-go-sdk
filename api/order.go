package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
)

type OrderService struct {
	Client *Client
}

// Create creates a new order
func (s *OrderService) Create(ctx context.Context, order *models.Order) (*models.Order, error) {
	req, err := s.Client.Request().Do(ctx, http.MethodPost, "/orders", order, &models.Order{})
	if err != nil {
		return nil, err
	}

	var response models.Order
	if err := s.Client.Request().Do(ctx, http.MethodPost, "/orders", order, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Get retrieves an order by ID
func (s *OrderService) Get(ctx context.Context, orderID string) (*models.Order, error) {
	path := fmt.Sprintf("/orders/%s", orderID)
	req, err := s.Client.Request().Do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response models.Order
	if err := s.Client.Request().Do(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
