package models

type Location struct {
	Latitude float64 `json:"latitude" dynamo:"Latitude"`
	Longitude float64 `json:"longitude" dynamo:"Longitude"`
}
