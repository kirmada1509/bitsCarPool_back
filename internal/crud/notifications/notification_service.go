package notification

import (
	"bitsCarPool_back/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationService interface{
	CreateNotification(notification *models.NotificationDetails) (string, error)
}

type notificationService struct{
	db *mongo.Client
	database string
}

func NewNotificationService(db *mongo.Client, database string) NotificationService{
	return &notificationService{
		db: db,
		database: database,
	}
}