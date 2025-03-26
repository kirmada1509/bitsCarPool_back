package models
type Notification struct{
	User_id string `json:"user_id" bson:"user_id"`
	Sent []SentNotification `json:"sent_notifications" bson:"sent_notifications"`
	Received []SentNotification `json:"received_notifications" bson:"received_notifications"`
}
type SentNotification struct {
	From    string `json:"from" bson:"from"`
	To string `json:"to" bson:"to"`
	TripId  string `json:"trip_id" bson:"trip_id"`
	Notes   string `json:"notes" bson:"notes"`
	Status string  `json:"status" bson:"status"` // "accept, decline"
}

type ReceivedNotification struct {
	From    string `json:"from" bson:"from"`
	To string `json:"to" bson:"to"`
	TripId  string `json:"trip_id" bson:"trip_id"`
	Notes   string `json:"notes" bson:"notes"`
	Status string  `json:"status" bson:"status"` // "accept, decline"
}