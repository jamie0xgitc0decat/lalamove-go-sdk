package client

import (
	"fmt"
	"net/http"
	"time"

	api "github.com/jamie0xgitc0decat/lalamove-go-sdk/api"
	"github.com/jamie0xgitc0decat/lalamove-go-sdk/internal/auth"
	"github.com/jamie0xgitc0decat/lalamove-go-sdk/internal/request"
)

const (
	defaultTimeout = 30 * time.Second
)

// Config holds the configuration for the Lalamove client
type Config struct {
	APIKey      string
	APISecret   string
	Environment Environment
	Market      string
	Timeout     time.Duration
	HTTPClient  *http.Client
}

// Client represents the Lalamove API client
type Client struct {
	// HTTP client used to communicate with the API
	httpClient *http.Client

	// Base URL for API requests
	baseURL string

	// Authentication credentials
	apiKey    string
	apiSecret string

	// Environment (sandbox or production)
	environment Environment

	// Service endpoints
	Orders     *api.OrderService
	Quotations *api.QuotationService
	Markets    *api.MarketService
	Drivers    *api.DriverService

	// Request handler
	request *request.Request

	// Market
	market string
}

// NewClient creates a new Lalamove API client with the given configuration
func NewClient(config Config) (*Client, error) {
	// Validate API credentials
	if err := auth.ValidateAPIKey(config.APIKey); err != nil {
		return nil, err
	}
	if err := auth.ValidateAPISecret(config.APISecret); err != nil {
		return nil, err
	}

	if config.Timeout == 0 {
		config.Timeout = defaultTimeout
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	if config.Environment == "" {
		config.Environment = Sandbox
	}

	if config.Market == "" {
		return nil, fmt.Errorf("market is required")
	}

	c := &Client{
		httpClient:  config.HTTPClient,
		baseURL:     config.Environment.GetBaseURL(),
		apiKey:      config.APIKey,
		apiSecret:   config.APISecret,
		environment: config.Environment,
		market:      config.Market,
	}

	// Initialize request handler
	c.request = request.NewRequest(
		c.httpClient,
		auth.Credentials{
			APIKey:    c.apiKey,
			APISecret: c.apiSecret,
		},
		c.baseURL,
		c.market,
	)

	// Initialize services
	c.Orders = &api.OrderService{Client: c.Orders.Client}
	c.Quotations = &api.QuotationService{Client: c}
	c.Markets = &api.MarketService{Client: c}
	c.Drivers = &api.DriverService{client: c}

	return c, nil
}

// IsProduction returns true if the client is configured for production
func (c *Client) IsProduction() bool {
	return c.environment == Production
}

// Request returns the request handler
func (c *Client) Request() *request.Request {
	return c.request
}
