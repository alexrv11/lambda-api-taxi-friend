package models

import "time"

type InputOrder struct {
	DriverID string `json:"driver_id"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Order struct {
	ID string `json:"id"`
	DriverID string `json:"driver_id"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status string `json:"status"`
	DateCreated time.Time `json:"date_created"`
	LastUpdated time.Time `json:"last_updated"`
}
