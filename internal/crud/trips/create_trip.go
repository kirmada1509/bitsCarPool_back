package trips

import (
	"bitsCarPool_back/internal/models"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (s *tripService) CreateTrip(trip *models.Trip) (string, error) {
	if trip == nil {
		return "", errors.New("invalid trip data")
	}
	if(trip.VehicleModel == ""){
		return "", errors.New("empty vehicle model")
	}
	fmt.Println("vehicle model, ", trip.VehicleModel)
	trip.CreatedAt = time.Now()
	trip.UpdatedAt = time.Now()

	collection := s.db.Database(s.database).Collection("trips")
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
