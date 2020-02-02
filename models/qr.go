package models

type Qr struct {
	Code string `json:"code"`
	Status string `json:"status"`
	Credit float64 `json:"credit"`
	DriverID string `json:"driver_id" dynamo:"DriverID"`
}
