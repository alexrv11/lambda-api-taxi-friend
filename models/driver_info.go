package models

//DriverInfo model
type DriverInfo struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Phone       string  `json:"phone"`
	Credit      float32 `json:"credit"`
	Status      string  `json:"status"`
	CarIdentity string  `json:"car_identity"`
}
