package main

import (
	"context"
	"log"
	"time"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/client"
	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
)

func main() {
	// Create a sandbox client
	sandboxClient, err := client.NewClient(client.Config{
		APIKey:      "pk_test_your-api-key",
		APISecret:   "sk_test_your-api-secret",
		Environment: client.Sandbox,
		Market:      "TH", // Thailand market
		Timeout:     60 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a production client
	prodClient := client.NewClient(client.Config{
		APIKey:      "your-prod-api-key",
		APISecret:   "your-prod-api-secret",
		Environment: client.Production,
		Timeout:     60 * time.Second,
	})

	// Example using sandbox client
	order, err := sandboxClient.Orders.Create(context.Background(), &models.Order{
		ServiceType: "MOTORCYCLE",
		Stops: []models.Stop{
			{
				Address: "123 Main St",
				Coordinates: models.Coordinates{
					Latitude:  1.234,
					Longitude: 4.567,
				},
				Contact: models.Contact{
					Name:        "John Doe",
					PhoneNumber: "+1234567890",
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("Failed to create order in sandbox: %v", err)
	}

	log.Printf("Created order in sandbox: %s", order.ID)
}
