package interfaces

import "github.com/jamie0xgitc0decat/lalamove-go-sdk/internal/request"

// Client defines the interface that clients must implement
type Client interface {
	Request() *request.Request
	IsProduction() bool
	Quotations() QuotationService
	Orders() OrderService
}
