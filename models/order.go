package models

import "time"

type Order struct {
	ID            string    `json:"id"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	QuotationID   string    `json:"quotation_id"`
	ServiceType   string    `json:"service_type"`
	SpecialRequests []string  `json:"special_requests"`
	
	Stops []Stop `json:"stops"`
}

type Stop struct {
	Coordinates Coordinates `json:"coordinates"`
	Address     string      `json:"address"`
	Contact     Contact     `json:"contact"`
}

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
} 