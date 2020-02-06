package models

//InputOrder model
type InputOrder struct {
	DriverID  string  `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//Order model
type Order struct {
	ID          string  `json:"id"`
	DriverID    string  `json:"driver_id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Status      string  `json:"status"`
	DateCreated string  `json:"date_created"`
	LastUpdated string  `json:"last_updated"`
}
