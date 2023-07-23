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

var fileCollection *mongo.Collection

func fileService(database *mongo.Database) {
	fileCollection = database.Collection("file")
}

func getFilesByBsonDocument(d primitive.D) ([]model.StorageFile, error) {
	var files []model.StorageFile

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := fileCollection.Find(ctx, d)
	if err != nil {
		return []model.StorageFile{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var file model.StorageFile
		cursor.Decode(&file)
		files = append(files, file)
	}

	if err := cursor.Err(); err != nil {
		return []model.StorageFile{}, err
	}

	return files, nil
}

func GetFiles() ([]model.StorageFile, error) {
	return getFilesByBsonDocument(bson.D{})
}

func GetFileListByMeeting(meeting string) ([]model.StorageFile, error) {
	return getFilesByBsonDocument(bson.D{{"meeting", meeting}, {"in_file_list", true}})
}

func GetFileById(id primitive.ObjectID) (model.StorageFile, error) {
	files, err := getFilesByBsonDocument(bson.D{{"_id", id}})
	if err != nil {
		return model.StorageFile{}, err
	}

	if len(files) > 0 {
		return files[0], nil
	}

	return model.StorageFile{}, errors.New("no entry with given id found")
}

func GetFileByNameAndMeeting(name string, meeting string) (model.StorageFile, error) {
	files, err := getFilesByBsonDocument(bson.D{{"name", name}, {"meeting", meeting}})
	if err != nil {
		return model.StorageFile{}, err
	}

	if len(files) > 0 {
		return files[0], nil
	}

	return model.StorageFile{}, errors.New("no entry with given name and meeting found")
}

func RemoveFileById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := fileCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func AddFile(file model.StorageFile) (model.StorageFile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	file.AddedAt = time.Now()
	file.UpdatedAt = time.Now()

	r, err := fileCollection.InsertOne(ctx, file)
	if err != nil {
		return model.StorageFile{}, err
	}

	return GetFileById(r.InsertedID.(primitive.ObjectID))
}

func UpdateFile(file model.StorageFile) (model.StorageFile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	file.UpdatedAt = time.Now()

	_, err := fileCollection.ReplaceOne(ctx, bson.D{{"_id", file.Identifier}}, file)
	if err != nil {
		return model.StorageFile{}, err
	}

	return GetFileById(file.Identifier)
}
