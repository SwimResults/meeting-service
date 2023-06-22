package service

import (
	"context"
	"errors"
	"github.com/swimresults/meeting-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var meetingCollection *mongo.Collection

func meetingService(database *mongo.Database) {
	meetingCollection = database.Collection("meeting")
}

func getMeetingsByBsonDocument(d primitive.D) ([]model.Meeting, error) {
	var meetings []model.Meeting

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := meetingCollection.Find(ctx, d)
	if err != nil {
		return []model.Meeting{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var meeting model.Meeting
		cursor.Decode(&meeting)

		meeting.Series, _ = GetMeetingSeriesById(meeting.SeriesId)

		if !meeting.LocationId.IsZero() {
			meeting.Location, _ = GetLocationById(meeting.LocationId)
		} else {
			meeting.Location, _ = GetLocationById(meeting.Series.LocationId)
		}

		if meeting.Organizer.IsZero() {
			meeting.Organizer = meeting.Series.Organizer
		}

		meetings = append(meetings, meeting)
	}

	if err := cursor.Err(); err != nil {
		return []model.Meeting{}, err
	}

	return meetings, nil
}

func GetMeetings() ([]model.Meeting, error) {
	return getMeetingsByBsonDocument(bson.D{})
}

func GetMeetingById(id primitive.ObjectID) (model.Meeting, error) {
	meetings, err := getMeetingsByBsonDocument(bson.D{{"_id", id}})
	if err != nil {
		return model.Meeting{}, err
	}

	if len(meetings) > 0 {
		return meetings[0], nil
	}

	return model.Meeting{}, errors.New("no entry with given id found")
}

func GetMeetingByMeetId(id string) (model.Meeting, error) {
	meetings, err := getMeetingsByBsonDocument(bson.D{{"meet_id", id}})
	if err != nil {
		return model.Meeting{}, err
	}

	if len(meetings) > 0 {
		return meetings[0], nil
	}

	return model.Meeting{}, errors.New("no entry with given meet_id found")
}

func GetMeetingsWithDateBetween(dateStart time.Time, dateEnd time.Time) ([]model.Meeting, error) {
	var result []model.Meeting
	meetings, err := getMeetingsByBsonDocument(bson.D{})
	if err != nil {
		return []model.Meeting{}, err
	}
	for _, meeting := range meetings {
		if meeting.DateStart.Second() >= dateStart.Second() && meeting.DateEnd.Second() <= dateEnd.Second() {
			result = append(result, meeting)
		}
	}
	return result, nil
}

func RemoveMeetingById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := meetingCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func AddMeeting(meeting model.Meeting) (model.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	meeting.SeriesId = meeting.Series.Identifier

	if !meeting.Location.Identifier.IsZero() {
		meeting.LocationId = meeting.Location.Identifier
	}

	meeting.AddedAt = time.Now()
	meeting.UpdatedAt = time.Now()

	r, err := meetingCollection.InsertOne(ctx, meeting)
	if err != nil {
		return model.Meeting{}, err
	}

	return GetMeetingById(r.InsertedID.(primitive.ObjectID))
}

func UpdateMeeting(meeting model.Meeting) (model.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	meeting.SeriesId = meeting.Series.Identifier

	if !meeting.Location.Identifier.IsZero() {
		meeting.LocationId = meeting.Location.Identifier
	}

	meeting.UpdatedAt = time.Now()

	_, err := meetingCollection.ReplaceOne(ctx, bson.D{{"_id", meeting.Identifier}}, meeting)
	if err != nil {
		return model.Meeting{}, err
	}

	return GetMeetingById(meeting.Identifier)
}
