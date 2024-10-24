package service

import (
	"context"
	"errors"
	"github.com/swimresults/meeting-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var incidentCollection *mongo.Collection

func incidentService(database *mongo.Database) {
	incidentCollection = database.Collection("incident")
}

func getIncidentsByBsonDocument(d primitive.D) ([]model.Incident, error) {
	var incidents []model.Incident

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryOptions := options.FindOptions{}
	queryOptions.SetSort(bson.D{{"start", 1}})

	cursor, err := incidentCollection.Find(ctx, d, &queryOptions)
	if err != nil {
		return []model.Incident{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var incident model.Incident
		cursor.Decode(&incident)
		incidents = append(incidents, incident)
	}

	if err := cursor.Err(); err != nil {
		return []model.Incident{}, err
	}

	return incidents, nil
}

func GetIncidentsByMeeting(meeting string) ([]model.Incident, error) {
	return getIncidentsByBsonDocument(bson.D{{"meeting", meeting}})
}

func GetIncidentById(id primitive.ObjectID) (model.Incident, error) {
	incidents, err := getIncidentsByBsonDocument(bson.D{{"_id", id}})
	if err != nil {
		return model.Incident{}, err
	}

	if len(incidents) > 0 {
		return incidents[0], nil
	}

	return model.Incident{}, errors.New("no entry with given id found")
}

func RemoveIncidentById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := incidentCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func AddIncident(incident model.Incident) (model.Incident, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	incident.AddedAt = time.Now()
	incident.UpdatedAt = time.Now()

	r, err := incidentCollection.InsertOne(ctx, incident)
	if err != nil {
		return model.Incident{}, err
	}

	return GetIncidentById(r.InsertedID.(primitive.ObjectID))
}

func UpdateIncident(incident model.Incident) (model.Incident, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	incident.UpdatedAt = time.Now()

	_, err := incidentCollection.ReplaceOne(ctx, bson.D{{"_id", incident.Identifier}}, incident)
	if err != nil {
		return model.Incident{}, err
	}

	return GetIncidentById(incident.Identifier)
}

func UpdateIncidentDateByMeeting(meeting string, t time.Time, updateTimeZone bool) ([]model.Incident, error) {
	incidents, err := GetIncidentsByMeeting(meeting)

	if err != nil {
		return []model.Incident{}, err
	}

	var savedIncidents []model.Incident

	for _, incident := range incidents {
		t1 := incident.Start
		t2 := incident.End

		var tz1 *time.Location
		var tz2 *time.Location
		if updateTimeZone {
			tz1 = t.Location()
			tz2 = t.Location()
		} else {
			tz1 = t1.Location()
			tz2 = t2.Location()
		}

		incident.End = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), tz1)
		incident.Start = time.Date(t2.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute(), t2.Second(), t2.Nanosecond(), tz2)

		saved, err := UpdateIncident(incident)
		if err != nil {
			return []model.Incident{}, err
		}

		savedIncidents = append(savedIncidents, saved)
	}

	return savedIncidents, nil
}
