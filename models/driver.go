package models

//Driver manages the driver information
type Driver struct {
	ID string `json:"id" dynamo:"Id,hash"`
	Name string `json:"name" dynamo:"Name"`
	FrontCarPhoto string `json:"frontCarPhoto" dynamo:"FrontCarPhoto"`
	BackCarPhoto string `json:"backCarPhoto" dynamo:"BackCarPhoto"`
	SideCarPhoto string `json:"sideCarPhoto" dynamo:"SideCarPhoto"`
	FrontLicensePhoto string `json:"frontLicensePhoto" dynamo:"FrontLicensePhoto"`
	BackLicensePhoto string `json:"backLicensePhoto" dynamo:"BackLicensePhoto"`
	Latitude float64 `json:"latitude" dynamo:"Latitude"`
	Longitude float64 `json:"longitude" dynamo:"Longitude"`
	Phone string `json:"phone" dynamo:"Phone"`
	Credit float32 `json:"credit" dynamo:"Credit"`
	Status string `json:"status" dynamo:"Status"`
	Direction int `json:"direction" dynamo:"Direction"`
	Password string `json:"password" dynamo:"Password"`
	CarIdentity string `json:"carIdentity" dynamo:"CarIdentity"`
}