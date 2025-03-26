package models
type NotificationModel struct {
	From    string `json:"from" bson:"from"`
	To      string `json:"to" bson:"to"`
	TripId  string `json:"trip_id" bson:"trip_id"`
	Notes   string `json:"notes" bson:"notes"`
	Type    string `json:"type" bson:"type"` // "request", "accept", "decline"
}

type UpdateNotificationModel struct {
	NotificationId string `json:"notification_id" bson:"notification_id"`
	Action         string `json:"action" bson:"action"` // Example: "read", "deleted", "updated"
}
