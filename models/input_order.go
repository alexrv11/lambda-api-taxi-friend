package models

import "time"

type InputOrder struct {
	ID string `json:"id" dynamo:"Id,hash"`
	DriverID string `json:"driver_id" dynamo:"DriverId"`
	Latitude float64 `json:"latitude" dynamo:"Latitude"`
	Longitude float64 `json:"longitude" dynamo:"Longitude"`
	Status string `json:"status" dynamo:"Status"`
	DateCreated time.Time `json:"date_created" dynamo:"DateCreated"`
	LastUpdated time.Time `json:"last_updated" dynamo:"LastUpdated"`
}
