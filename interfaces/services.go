package interfaces

import (
	"context"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
)

type QuotationService interface {
	Create(ctx context.Context, req *models.QuotationRequest) (*models.QuotationData, error)
	Get(ctx context.Context, quotationID string) (*models.QuotationData, error)
}

type OrderService interface {
	Create(ctx context.Context, order *models.Order) (*models.Order, error)
	Get(ctx context.Context, orderID string) (*models.Order, error)
}
