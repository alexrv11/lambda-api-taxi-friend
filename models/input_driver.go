package models

type InputDriver struct {
	Id string `json:"id" dynamo:"Id,hash"`
	Name string `json:"name" dynamo:"Name"`
	FrontCarPhoto string `json:"front_car_photo" dynamo:"FrontCarPhoto"`
	BackCarPhoto string `json:"back_car_photo" dynamo:"BackCarPhoto"`
	SideCarPhoto string `json:"side_car_photo" dynamo:"SideCarPhoto"`
	FrontLicensePhoto string `json:"front_license_photo" dynamo:"FrontLicensePhoto"`
	BackLicensePhoto string `json:"back_license_photo" dynamo:"BackLicensePhoto"`
	Latitude float32 `json:"latitude" dynamo:"Latitude"`
	Longitude float32 `json:"longitude" dynamo:"Longitude"`
	Phone string `json:"phone" dynamo:"Phone"`
	Credit float32 `json:"credit" dynamo:"Credit"`
	Status string `json:"status" dynamo:"Status"`
	Direction int `json:"direction" dynamo:"Direction"`
	Password string `json:"password" dynamo:"Password"`
	CarIdentity string `json:"car_identity" dynamo:"CarIdentity"`
}