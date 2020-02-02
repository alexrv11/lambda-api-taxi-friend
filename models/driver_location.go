package models

type DriverLocation struct {
	ID string `json:"id"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status string `json:"status"`
	Direction int `json:"direction"`
}
