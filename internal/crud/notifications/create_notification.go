package notification

import (
	"bitsCarPool_back/internal/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *notificationService) CreateNotification(notification *models.Notification) (string, error) {
	if notification == nil {
		return "", errors.New("invalid data")
	}


	notification.NotificationID = primitive.NewObjectID().Hex()

	collection := s.db.Database(s.database).Collection("users")

	// Add to recipient's notifications_received
	receivedUpdate := bson.M{
		"$push": bson.M{
			"notifications_received": notification,
		},
	}

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": notification.To},
		receivedUpdate,
	)
	if err != nil {
		return "", fmt.Errorf("failed to update notifications_received for user %s: %v", notification.To, err)
	}

	// Add to sender's notifications_sent
	sentUpdate := bson.M{
		"$push": bson.M{
			"notifications_sent": notification,
		},
	}

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": notification.From},
		sentUpdate,
	)
	if err != nil {
		return "", fmt.Errorf("failed to update notifications_sent for user %s: %v", notification.From, err)
	}

	return notification.NotificationID, nil
}
