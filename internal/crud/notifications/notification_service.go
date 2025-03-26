package notification

import (
	"bitsCarPool_back/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationService interface{
	GetNotifications(user_id string) ([]models.NotificationModel, error)
	CreateNotification(notification *models.NotificationModel) (string, error)
	UpdateNotification(updateNotificationModel *models.UpdateNotificationModel) (string, error)
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

func (s *notificationService) GetNotifications(user_id string) ([]models.NotificationModel, error){
	var n []models.NotificationModel
	return n, nil
}
