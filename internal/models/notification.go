package models
type Notification struct {
	User_id    string         `json:"user_id" bson:"user_id"`
	Sent       []NotificationDetails `json:"sent_notifications" bson:"sent_notifications"`
	Received   []NotificationDetails `json:"received_notifications" bson:"received_notifications"`
}

type NotificationDetails struct {
	From   string `json:"from" bson:"from"`
	To     string `json:"to" bson:"to"`
	TripId string `json:"trip_id" bson:"trip_id"`
	Notes  string `json:"notes" bson:"notes"`
	Status string `json:"status" bson:"status"` // "accept", "decline"
}
