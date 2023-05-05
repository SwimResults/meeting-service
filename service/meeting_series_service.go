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

func meetingSeriesService(database *mongo.Database) {
	collection = database.Collection("meeting_series")
}

func getMeetingSeriesByBsonDocument(d primitive.D) ([]model.MeetingSeries, error) {
	var meetings []model.MeetingSeries

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, d)
	if err != nil {
		return []model.MeetingSeries{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var meeting model.MeetingSeries
		cursor.Decode(&meeting)
		if !meeting.LocationId.IsZero() {
			meeting.Location, _ = GetLocationById(meeting.LocationId)
		}
		meetings = append(meetings, meeting)
	}

	if err := cursor.Err(); err != nil {
		return []model.MeetingSeries{}, err
	}

	return meetings, nil
}

func GetMeetingSeries() ([]model.MeetingSeries, error) {
	return getMeetingSeriesByBsonDocument(bson.D{})
}

func GetMeetingSeriesById(id primitive.ObjectID) (model.MeetingSeries, error) {
	meetings, err := getMeetingSeriesByBsonDocument(bson.D{{"_id", id}})
	if err != nil {
		return model.MeetingSeries{}, err
	}

	if len(meetings) > 0 {
		return meetings[0], nil
	}

	return model.MeetingSeries{}, errors.New("no entry with given id found")
}

func RemoveMeetingSeriesById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func AddMeetingSeries(meeting model.MeetingSeries) (model.MeetingSeries, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if !meeting.Location.Identifier.IsZero() {
		meeting.LocationId = meeting.Location.Identifier
	}

	meeting.AddedAt = time.Now()
	meeting.UpdatedAt = time.Now()

	r, err := collection.InsertOne(ctx, meeting)
	if err != nil {
		return model.MeetingSeries{}, err
	}

	return GetMeetingSeriesById(r.InsertedID.(primitive.ObjectID))
}

func UpdateMeetingSeries(meeting model.MeetingSeries) (model.MeetingSeries, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	meeting.UpdatedAt = time.Now()

	_, err := collection.ReplaceOne(ctx, bson.D{{"_id", meeting.Identifier}}, meeting)
	if err != nil {
		return model.MeetingSeries{}, err
	}

	return GetMeetingSeriesById(meeting.Identifier)
}
