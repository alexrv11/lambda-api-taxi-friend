package models

//Qr model
type Qr struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Credit float64 `json:"credit"`
	DriverID string `json:"driver_id"`
}

//InputQrDriver model
type InputQrDriver struct {
	DriverID string `json:"driver_id"`
}
