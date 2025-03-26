package models

type Notification struct {
	NotificationID string`json:"notification_id" bson:"notification_id,omitempty"`
	From    string `json:"from" bson:"from"`
	To      string `json:"to" bson:"to"`
	TripId  string `json:"trip_id" bson:"trip_id"`
	Notes   string `json:"notes" bson:"notes"`
	Status string  `json:"status" bson:"status"` // "accept, decline"
}