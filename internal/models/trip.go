package models

import "time"

type Trip struct {
	ID                string `json:"id" bson:"_id,omitempty"`
	From             string    `json:"from" bson:"from"`
	To               string    `json:"to" bson:"to"`
	DepartureTime    time.Time `json:"departure_time" bson:"departure_time"`
	FlexibilityWindow int      `json:"flexibility_window" bson:"flexibility_window"`
	Vehicle          Vehicle   `json:"vehicle" bson:"vehicle"`
	SeatsAvailable   int       `json:"seats_available" bson:"seats_available"`
	TotalFare        float64   `json:"total_fare" bson:"total_fare"`
	Notes            string    `json:"notes,omitempty" bson:"notes,omitempty"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
	Creator          Student   `json:"creator" bson:"creator"`
}
