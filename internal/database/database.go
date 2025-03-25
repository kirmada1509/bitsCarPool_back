package database

import (
	"bitsCarPool_back/internal/models"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	CreateTrip(trip *models.Trip) (string, error)

}

type service struct {
	db *mongo.Client
}

var (
	host = os.Getenv("BLUEPRINT_DB_HOST")
	port = os.Getenv("BLUEPRINT_DB_PORT")
	database = os.Getenv("BLUEPRINT_DB_DATABASE")
)

func New() Service {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))

	if err != nil {
		log.Fatal(err)

	}
	return &service{
		db: client,
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


func(s *service) CreateTrip(trip *models.Trip) (string, error) {
	if trip == nil {
		return "", errors.New("invalid trip data")
	}
	trip.CreatedAt = time.Now()
	trip.UpdatedAt = time.Now()

	collection := s.db.Database(database).Collection("trips")
	res, err := collection.InsertOne(context.Background(), trip)
	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to retrieve trip ID")
	}
	return id.Hex(), nil
}
