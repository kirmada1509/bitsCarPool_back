package notification

import (
	"bitsCarPool_back/internal/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *notificationService) CreateNotification(notification *models.NotificationDetails) (string, error) {
	collection := s.db.Database(s.database).Collection("notifications")

	filter := bson.M{"user_id": notification.From}
	result := collection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$push": bson.M{"sent_notifications": notification}})
	if result.Err() != nil {
		return "", result.Err()
	}

	filter = bson.M{"user_id": notification.To}
	result = collection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$push": bson.M{"received_notifications": notification}})
	if result.Err() != nil {
		return "", result.Err()
	}

	return fmt.Sprint("notification sent from ", notification.From, " to ", notification.To), nil
}

func (s *notificationService) GetNotifications(user *models.User) (models.Notification, error) {
	collection := s.db.Database(s.database).Collection("notifications")

	filter := bson.M{"user_id": user.ID}
	var result models.Notification

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Notification{User_id: user.ID}, err
	}

	return result, nil
}
