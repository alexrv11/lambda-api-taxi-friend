package models

//Driver manages the driver information
type Driver struct {
	ID string `json:"id"`
	Name string `json:"name"`
	FrontCarPhoto string `json:"front_car_photo"`
	BackCarPhoto string `json:"back_car_photo"`
	SideCarPhoto string `json:"side_car_photo"`
	FrontLicensePhoto string `json:"front_license_photo"`
	BackLicensePhoto string `json:"back_license_photo"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Phone string `json:"phone"`
	Credit float32 `json:"credit"`
	Status string `json:"status"`
	Direction int `json:"direction"`
	Password string `json:"password"`
	CarIdentity string `json:"car_identity"`
}