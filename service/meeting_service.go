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

var collection *mongo.Collection

func meetingService(database *mongo.Database) {
	collection = database.Collection("meeting")
}

func GetMeetings() ([]model.Meeting, error) {
	var meetings []model.Meeting

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []model.Meeting{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var meeting model.Meeting
		cursor.Decode(&meeting)
		meetings = append(meetings, meeting)
	}

	if err := cursor.Err(); err != nil {
		return []model.Meeting{}, err
	}

	return meetings, nil
}

func GetMeetingById(id primitive.ObjectID) (model.Meeting, error) {
	var meeting model.Meeting

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{{"_id", id}})
	if err != nil {
		return model.Meeting{}, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		cursor.Decode(&meeting)
		return meeting, nil
	}

	return model.Meeting{}, errors.New("no entry with given id found")
}

func RemoveMeetingById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func AddMeeting(meeting model.Meeting) (model.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := collection.InsertOne(ctx, meeting)
	if err != nil {
		return model.Meeting{}, err
	}

	return GetMeetingById(r.InsertedID.(primitive.ObjectID))
}

func UpdateMeeting(meeting model.Meeting) (model.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.ReplaceOne(ctx, bson.D{{"_id", meeting.Identifier}}, meeting)
	if err != nil {
		return model.Meeting{}, err
	}

	return GetMeetingById(meeting.Identifier)
}
