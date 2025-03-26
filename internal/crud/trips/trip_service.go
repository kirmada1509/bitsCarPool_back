package trip

import (
	"bitsCarPool_back/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TripService interface {
	CreateTrip(trip *models.Trip) (string, error)
	SearchTrips(search models.TripSearch) ([]models.Trip, error)
}

type tripService struct {
	db       *mongo.Client
	database string
}

func NewTripService(db *mongo.Client, database string) TripService {
	return &tripService{
		db:       db,
		database: database,
	}
}
