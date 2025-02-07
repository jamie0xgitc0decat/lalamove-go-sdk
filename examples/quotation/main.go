package main

import (
	"context"
	"log"
	"time"

	"github.com/jamie0xgitc0decat/lalamove-go-sdk/client"
	"github.com/jamie0xgitc0decat/lalamove-go-sdk/models"
)

func main() {
	// Create client
	lalamove, err := client.NewClient(client.Config{
		APIKey:      "pk_test_your-api-key",
		APISecret:   "sk_test_your-api-secret",
		Environment: client.Sandbox,
		Market:      "HK",
		Timeout:     60 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create quotation request
	req := &models.QuotationRequest{
		ServiceType: "MOTORCYCLE",
		Language:    "en_HK",
		Stops: []models.Stop{
			{
				Address: "Innocentre, 72 Tat Chee Ave, Kowloon Tong",
				Coordinates: models.Coordinates{
					Latitude:  22.33547351186244,
					Longitude: 114.17615807116502,
				},
			},
			{
				Address: "Statue Square, Des Voeux Rd Central, Central",
				Coordinates: models.Coordinates{
					Latitude:  22.28129462633954,
					Longitude: 114.15986100706951,
				},
			},
		},
		Item: &models.Item{
			Quantity:             "3",
			Weight:               "LESS_THAN_3KG",
			Categories:           []string{"FOOD_DELIVERY", "OFFICE_ITEM"},
			HandlingInstructions: []string{"KEEP_UPRIGHT"},
		},
		IsRouteOptimized: true,
	}

	// Get quotation
	quotation, err := lalamove.Quotations.Create(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create quotation: %v", err)
	}

	log.Printf("Created quotation: %s", quotation.QuotationID)
	log.Printf("Total price: %s %s", quotation.PriceBreakdown.Total, quotation.PriceBreakdown.Currency)
	log.Printf("Distance: %f %s", quotation.Distance.Value, quotation.Distance.Unit)
}

func getQuotationExample() {
	// Create client
	lalamove, err := client.NewClient(client.Config{
		APIKey:      "pk_test_your-api-key",
		APISecret:   "sk_test_your-api-secret",
		Environment: client.Sandbox,
		Market:      "HK",
		Timeout:     60 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get quotation details
	quotationID := "your-quotation-id"
	quotation, err := lalamove.Quotations.Get(context.Background(), quotationID)
	if err != nil {
		log.Fatalf("Failed to get quotation details: %v", err)
	}

	// Print quotation details
	log.Printf("Quotation ID: %s", quotation.QuotationID)
	log.Printf("Service Type: %s", quotation.ServiceType)
	log.Printf("Expires At: %s", quotation.ExpiresAt.Format(time.RFC3339))
	log.Printf("Total Price: %s %s", quotation.PriceBreakdown.Total, quotation.PriceBreakdown.Currency)

	// Print stops
	for i, stop := range quotation.Stops {
		log.Printf("Stop %d: %s (%.6f, %.6f)",
			i+1,
			stop.Address,
			stop.Coordinates.Latitude,
			stop.Coordinates.Longitude,
		)
	}
}
