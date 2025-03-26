package models

import "time"

type Trip struct {
	ID                string `json:"id" bson:"_id,omitempty"`
	Creator          Student   `json:"creator" bson:"creator"`
	From             string    `json:"from" bson:"from"`
	To               string    `json:"to" bson:"to"`
	DepartureTime    time.Time `json:"departure_time" bson:"departure_time"`
	FlexibilityWindow int      `json:"flexibility_window" bson:"flexibility_window"`
	TotalFare        float64   `json:"total_fare" bson:"total_fare"`
	VehicleModel        string  `json:"vehicle_model" bson:"vehicle_model"`
	Capacity     int     `json:"capacity" bson:"capacity"` 
	SeatsAvailable   int       `json:"seats_available" bson:"seats_available"`
	FuelType     string  `json:"fuel_type" bson:"fuel_type"` 
	AC           bool    `json:"ac" bson:"ac"` 
	Notes            string    `json:"notes,omitempty" bson:"notes,omitempty"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}


type TripSearch struct {
	From              *string    `json:"from" bson:"from"`
	To                *string    `json:"to" bson:"to"`
	DepartureTime     *time.Time `json:"departure_time" bson:"departure_time"`
	FlexibilityWindow *int      `json:"flexibility_window" bson:"flexibility_window"`
	SeatsAvailable    *int       `json:"seats_available" bson:"seats_available"`
	Capacity     int     `json:"capacity" bson:"capacity"` 
	TotalFare         *float64   `json:"total_fare" bson:"total_fare"`
	AC           bool    `json:"ac" bson:"ac"` 
	FuelType     string  `json:"fuel_type" bson:"fuel_type"` 
}