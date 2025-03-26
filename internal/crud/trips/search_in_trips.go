package trips

import (
	"bitsCarPool_back/internal/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *tripService) SearchTrips(search models.TripSearch) ([]models.Trip, error) {
	if search.From == nil || search.To == nil {
		return nil, errors.New("from and to locations are required")
	}

	filter := bson.M{
		"from": *search.From,
		"to":   *search.To,
	}

	if search.DepartureTime != nil && search.FlexibilityWindow != nil {
		filter["departure_time"] = bson.M{
			"$gte": search.DepartureTime.Add(-time.Duration(*search.FlexibilityWindow) * time.Minute),
			"$lte": search.DepartureTime.Add(time.Duration(*search.FlexibilityWindow) * time.Minute),
		}
	}
	if search.Capacity > 0 {
		filter["vehicle.capacity"] = bson.M{"$gte": search.Capacity}
	}
	if search.FuelType != "" {
		filter["vehicle.fuel_type"] = search.FuelType
	}
	if search.AC {
		filter["vehicle.ac"] = true
	}
	if search.SeatsAvailable != nil {
		filter["seats_available"] = bson.M{"$gte": *search.SeatsAvailable}
	}
	if search.TotalFare != nil {
		filter["total_fare"] = bson.M{"$lte": *search.TotalFare}
	}

	collection := s.db.Database(s.database).Collection("trips")
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var trips []models.Trip
	if err := cursor.All(context.Background(), &trips); err != nil {
		return nil, err
	}

	return trips, nil
}
