package models

//Driver manages the driver information
type Driver struct {
	ID string `json:"id"`
	Name string `json:"name"`
	FrontCarPhoto string `json:"frontCarPhoto"`
	BackCarPhoto string `json:"backCarPhoto"`
	SideCarPhoto string `json:"sideCarPhoto"`
	FrontLicensePhoto string `json:"frontLicensePhoto"`
	BackLicensePhoto string `json:"backLicensePhoto"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Phone string `json:"phone"`
	Credit float32 `json:"credit"`
	Status string `json:"status"`
	Direction int `json:"direction"`
	Password string `json:"password"`
	CarIdentity string `json:"carIdentity"`
}