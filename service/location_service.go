package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sr-meeting/meeting-service/model"
	"time"
)

var locationCollection *mongo.Collection

func locationService(database *mongo.Database) {
	locationCollection = database.Collection("location")
}

func GetLocations() ([]model.Location, error) {
	var locations []model.Location

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := locationCollection.Find(ctx, bson.M{})
	if err != nil {
		return []model.Location{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var location model.Location
		cursor.Decode(&location)
		locations = append(locations, location)
	}

	if err := cursor.Err(); err != nil {
		return []model.Location{}, err
	}

	return locations, nil
}

func GetLocationById(id primitive.ObjectID) (model.Location, error) {
	var location model.Location

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := locationCollection.Find(ctx, bson.D{{"_id", id}})
	if err != nil {
		return model.Location{}, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		cursor.Decode(&location)
		return location, nil
	}

	return model.Location{}, errors.New("no entry with given id found")
}

func AddLocation(location model.Location) (model.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	location.AddedAt = time.Now()
	location.UpdatedAt = time.Now()

	r, err := locationCollection.InsertOne(ctx, location)
	if err != nil {
		return model.Location{}, err
	}

	return GetLocationById(r.InsertedID.(primitive.ObjectID))
}
