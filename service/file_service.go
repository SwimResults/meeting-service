package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/swimresults/meeting-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	queryOptions := options.FindOptions{}
	queryOptions.SetSort(bson.D{{"ordering", 1}})

	cursor, err := fileCollection.Find(ctx, d, &queryOptions)
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

	if file.Path != "" && file.Path[0] != '/' {
		file.Path = "/" + file.Path
	}

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

	if file.Path != "" && file.Path[0] != '/' {
		file.Path = "/" + file.Path
	}

	_, err := fileCollection.ReplaceOne(ctx, bson.D{{"_id", file.Identifier}}, file)
	if err != nil {
		return model.StorageFile{}, err
	}

	return GetFileById(file.Identifier)
}

func IncrementDownloads(path string) (bool, error) {
	if path == "" {
		return false, nil
	}
	if path[0] != '/' {
		path = "/" + path
	}
	fmt.Printf("Incrementing downloads on file: '%s'\n", path)
	files, err := getFilesByBsonDocument(bson.D{{"path", path}})
	if err != nil {
		return false, err
	}

	if len(files) < 1 {
		return false, nil
	}

	var file model.StorageFile
	file = files[0]
	file.Downloads++

	_, err = UpdateFile(file)
	if err != nil {
		return true, err
	}

	return true, nil
}
