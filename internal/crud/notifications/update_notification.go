package notification

import (
	"bitsCarPool_back/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *notificationService) UpdateNotification(updateNotificationModel *models.UpdateNotificationModel) (string, error){
	
	users_collection := s.db.Database(s.database).Collection("users")
	
	filter := bson.M{"notifications_received.notification_id": updateNotificationModel.NotificationId}
	update := bson.M{"$set": bson.D{{Key: "status", Value: updateNotificationModel.Action}}}
	
	_, err := users_collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return "", err
	}

	return  "updated notification", nil
}