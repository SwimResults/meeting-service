package service

import (
	"context"
	"errors"
	"github.com/swimresults/meeting-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

var ageGroupCollection *mongo.Collection

func ageGroupService(database *mongo.Database) {
	ageGroupCollection = database.Collection("age_group")
}

var ageGroupNotFoundError = "age group not found"

func getAgeGroupsByBsonDocument(d primitive.D) ([]model.AgeGroup, error) {
	var ageGroups []model.AgeGroup

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryOptions := options.FindOptions{}
	queryOptions.SetSort(bson.D{{"min_age", 1}})

	cursor, err := ageGroupCollection.Find(ctx, d, &queryOptions)
	if err != nil {
		return []model.AgeGroup{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var ageGroup model.AgeGroup
		cursor.Decode(&ageGroup)
		ageGroups = append(ageGroups, ageGroup)
	}

	if err := cursor.Err(); err != nil {
		return []model.AgeGroup{}, err
	}

	return ageGroups, nil
}

func getAgeGroupByBsonDocument(d primitive.D) (model.AgeGroup, error) {
	groups, err := getAgeGroupsByBsonDocument(d)
	if err != nil {
		return model.AgeGroup{}, err
	}

	if len(groups) <= 0 {
		return model.AgeGroup{}, errors.New(ageGroupNotFoundError)
	}

	return groups[0], nil
}

func GetAgeGroups() ([]model.AgeGroup, error) {
	return getAgeGroupsByBsonDocument(bson.D{})
}

func GetAgeGroupsByMeeting(meeting string) ([]model.AgeGroup, error) {
	return getAgeGroupsByBsonDocument(bson.D{{"meeting", meeting}})
}

func GetAgeGroupsByMeetingAndEvent(meeting string, event int) ([]model.AgeGroup, error) {
	return getAgeGroupsByBsonDocument(bson.D{{"meeting", meeting}, {"event", event}})
}

func GetAgeGroupByMeetingAndEventAndAgesAndGender(meeting string, event int, minAge string, maxAge string, gender string) (model.AgeGroup, error) {
	return getAgeGroupByBsonDocument(bson.D{{"meeting", meeting}, {"event", event}, {"min_age", minAge}, {"max_age", maxAge}, {"gender", gender}})
}

func GetAgeGroupById(id primitive.ObjectID) (model.AgeGroup, error) {
	ageGroups, err := getAgeGroupsByBsonDocument(bson.D{{"_id", id}})
	if err != nil {
		return model.AgeGroup{}, err
	}

	if len(ageGroups) > 0 {
		return ageGroups[0], nil
	}

	return model.AgeGroup{}, errors.New("no entry with given id found")
}

func RemoveAgeGroupById(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := ageGroupCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func ImportAgeGroup(group model.AgeGroup) (*model.AgeGroup, bool, error) {
	existing, err := GetAgeGroupByMeetingAndEventAndAgesAndGender(group.Meeting, group.Event, group.MinAge, group.MaxAge, group.Gender)
	if err != nil {
		if err.Error() != ageGroupNotFoundError {
			return nil, false, err
		}

		newGroup, err2 := AddAgeGroup(group)
		if err2 != nil {
			return nil, false, err2
		}
		return &newGroup, false, nil
	}

	group.Identifier = existing.Identifier

	newGroup, err := UpdateAgeGroup(group)
	if err != nil {
		return nil, false, err
	}
	return &newGroup, false, nil
}

func AddAgeGroup(ageGroup model.AgeGroup) (model.AgeGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ageGroup.AddedAt = time.Now()
	ageGroup.UpdatedAt = time.Now()

	SetAgesForAgeGroup(&ageGroup)

	r, err := ageGroupCollection.InsertOne(ctx, ageGroup)
	if err != nil {
		return model.AgeGroup{}, err
	}

	return GetAgeGroupById(r.InsertedID.(primitive.ObjectID))
}

func UpdateAgeGroup(ageGroup model.AgeGroup) (model.AgeGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ageGroup.UpdatedAt = time.Now()

	SetAgesForAgeGroup(&ageGroup)

	_, err := ageGroupCollection.ReplaceOne(ctx, bson.D{{"_id", ageGroup.Identifier}}, ageGroup)
	if err != nil {
		return model.AgeGroup{}, err
	}

	return GetAgeGroupById(ageGroup.Identifier)
}

func SetAgesForAgeGroup(group *model.AgeGroup) {
	if !group.IsYear {
		return
	}

	min, _ := strconv.Atoi(group.MinAge)
	max, _ := strconv.Atoi(group.MaxAge)

	if min > max {
		a := min
		min = max
		max = a
	}

	for i := min; i <= max; i++ {
		if i < 1900 || i > 2050 {
			continue
		}
		group.Ages = append(group.Ages, strconv.Itoa(i))
	}
}
