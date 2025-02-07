package models

import "time"

// QuotationRequest represents the request body for creating a quotation
type QuotationRequest struct {
	ScheduleAt       *time.Time `json:"scheduleAt,omitempty"`
	ServiceType      string     `json:"serviceType"`
	SpecialRequests  []string   `json:"specialRequests,omitempty"`
	Language         string     `json:"language"`
	Stops            []Stop     `json:"stops"`
	Item            *Item      `json:"item,omitempty"`
	IsRouteOptimized bool       `json:"isRouteOptimized,omitempty"`
}

// QuotationResponse represents the response from the quotation API
type QuotationResponse struct {
	Data QuotationData `json:"data"`
}

// QuotationData represents the quotation details
type QuotationData struct {
	QuotationID      string           `json:"quotationId"`
	ScheduleAt       time.Time        `json:"scheduleAt"`
	ExpiresAt        time.Time        `json:"expiresAt"`
	ServiceType      string           `json:"serviceType"`
	SpecialRequests  []string         `json:"specialRequests"`
	Language         string           `json:"language"`
	Stops            []QuotationStop  `json:"stops"`
	IsRouteOptimized bool             `json:"isRouteOptimized"`
	PriceBreakdown   PriceBreakdown   `json:"priceBreakdown"`
	Item             *Item            `json:"item,omitempty"`
	Distance         Distance         `json:"distance"`
}

// QuotationStop represents a stop in the quotation
type QuotationStop struct {
	StopID      string      `json:"stopId"`
	Coordinates Coordinates `json:"coordinates"`
	Address     string      `json:"address"`
}

// PriceBreakdown represents the price breakdown of a quotation
type PriceBreakdown struct {
	Base                    string `json:"base"`
	SpecialRequests        string `json:"specialRequests"`
	VAT                     string `json:"vat"`
	TotalBeforeOptimization string `json:"totalBeforeOptimization"`
	TotalExcludePriorityFee string `json:"totalExcludePriorityFee"`
	Total                   string `json:"total"`
	Currency               string `json:"currency"`
}

// Distance represents the distance information
type Distance struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// Item represents the item details for delivery
type Item struct {
	Quantity             string   `json:"quantity,omitempty"`
	Weight               string   `json:"weight"`
	Categories           []string `json:"categories"`
	HandlingInstructions []string `json:"handlingInstructions,omitempty"`
} 