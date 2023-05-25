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

var eventCollection *mongo.Collection

func eventService(database *mongo.Database) {
	eventCollection = database.Collection("event")
}

func getEventsByBsonDocument(d primitive.D) ([]model.Event, error) {
	var events []model.Event

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := eventCollection.Find(ctx, d)
	if err != nil {
		return []model.Event{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var event model.Event
		cursor.Decode(&event)

		if !event.StyleId.IsZero() {
			event.Style, _ = GetStyleById(event.StyleId)
		}

		events = append(events, event)
	}

	if err := cursor.Err(); err != nil {
		return []model.Event{}, err
	}

	return events, nil
}

func GetEvents() ([]model.Event, error) {
	return getEventsByBsonDocument(bson.D{})
}

func GetEventById(id primitive.ObjectID) (model.Event, error) {
	events, err := getEventsByBsonDocument(bson.D{{"_id", id}})
	if err != nil {
		return model.Event{}, err
	}

	if len(events) > 0 {
		return events[0], nil
	}

	return model.Event{}, errors.New("no entry with given id found")
}

func GetEventByMeetId(id string) (model.Event, error) {
	events, err := getEventsByBsonDocument(bson.D{{"meeting", id}})
	if err != nil {
		return model.Event{}, err
	}

	if len(events) > 0 {
		return events[0], nil
	}

	return model.Event{}, errors.New("no entry with given meeting found")
}

func RemoveEventById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := eventCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func AddEvent(event model.Event) (model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if !event.Style.Identifier.IsZero() {
		event.StyleId = event.Style.Identifier
	}

	event.AddedAt = time.Now()
	event.UpdatedAt = time.Now()

	r, err := eventCollection.InsertOne(ctx, event)
	if err != nil {
		return model.Event{}, err
	}

	return GetEventById(r.InsertedID.(primitive.ObjectID))
}

func UpdateEvent(event model.Event) (model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if !event.Style.Identifier.IsZero() {
		event.StyleId = event.Style.Identifier
	}

	event.UpdatedAt = time.Now()

	_, err := eventCollection.ReplaceOne(ctx, bson.D{{"_id", event.Identifier}}, event)
	if err != nil {
		return model.Event{}, err
	}

	return GetEventById(event.Identifier)
}
