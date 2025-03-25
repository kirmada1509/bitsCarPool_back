package models

type Vehicle struct {
	ID           string  `json:"id" bson:"_id,omitempty"` 
	Model        string  `json:"model" bson:"model"`         
	Capacity     int     `json:"capacity" bson:"capacity"` 
	FuelType     string  `json:"fuel_type" bson:"fuel_type"` 
	AC           bool    `json:"ac" bson:"ac"` 
}
