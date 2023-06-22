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

var styleCollection *mongo.Collection

func styleService(database *mongo.Database) {
	styleCollection = database.Collection("style")
}

func getStyleByBsonDocument(d primitive.D) ([]model.Style, error) {
	var styles []model.Style

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := styleCollection.Find(ctx, d)
	if err != nil {
		return []model.Style{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var style model.Style
		cursor.Decode(&style)
		styles = append(styles, style)
	}

	if err := cursor.Err(); err != nil {
		return []model.Style{}, err
	}

	return styles, nil
}

func GetStyles() ([]model.Style, error) {
	return getStyleByBsonDocument(bson.D{})
}

func GetStyleById(id primitive.ObjectID) (model.Style, error) {
	styles, err := getStyleByBsonDocument(bson.D{{"_id", id}})
	if err != nil {
		return model.Style{}, err
	}

	if len(styles) > 0 {
		return styles[0], nil
	}

	return model.Style{}, errors.New("no entry with given id found")
}

func GetStyleByName(name string) (model.Style, error) {
	styles, err := getStyleByBsonDocument(bson.D{{"aliases", name}})
	if err != nil {
		return model.Style{}, err
	}

	if len(styles) > 0 {
		return styles[0], nil
	}

	return model.Style{}, errors.New("no entry with given name found")
}

func RemoveStyleById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := styleCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func AddStyle(style model.Style) (model.Style, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := styleCollection.InsertOne(ctx, style)
	if err != nil {
		return model.Style{}, err
	}

	return GetStyleById(r.InsertedID.(primitive.ObjectID))
}

func UpdateStyle(style model.Style) (model.Style, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := styleCollection.ReplaceOne(ctx, bson.D{{"_id", style.Identifier}}, style)
	if err != nil {
		return model.Style{}, err
	}

	return GetStyleById(style.Identifier)
}
