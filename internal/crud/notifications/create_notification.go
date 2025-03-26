package notification

import (
	"bitsCarPool_back/internal/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *notificationService) CreateNotification(notification *models.NotificationDetails) (string, error){
	// go to collection: notifications
	//access from and put in sent
	//access to and put in received

	collection := s.db.Database(s.database).Collection("notifications")

	filter := bson.M{"user_id" : notification.From}
	result := collection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$push": bson.M{"sent_notifications": notification}})
	if result.Err() != nil {
		return "", result.Err()
	}

	filter = bson.M{"user_id" : notification.To}
	result = collection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$push": bson.M{"received_notifications": notification}})
	if result.Err() != nil {
		return "", result.Err()
	}

	return fmt.Sprint("notification sent from ", notification.From, " to ", notification.To), nil
}
	