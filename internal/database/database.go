package database

import (
	"bitsCarPool_back/internal/crud/trips"
	"bitsCarPool_back/internal/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	CreateTrip(trip *models.Trip) (string, error)
}

type service struct {
	db    *mongo.Client
	trips trips.TripService
}

var (
	host     = os.Getenv("BLUEPRINT_DB_HOST")
	port     = os.Getenv("BLUEPRINT_DB_PORT")
	database = os.Getenv("BLUEPRINT_DB_DATABASE")
)

func New() Service {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))
	if err != nil {
		log.Fatal(err)
	}

	return &service{
		db:    client,
		trips: trips.NewTripService(client, database),
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) CreateTrip(trip *models.Trip) (string, error) {
	return s.trips.CreateTrip(trip)
}
