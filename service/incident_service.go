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
