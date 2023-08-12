package service

import (
	"context"
	"errors"
	"github.com/swimresults/meeting-service/dto"
	"github.com/swimresults/meeting-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
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

	queryOptions := options.FindOptions{}
	queryOptions.SetSort(bson.D{{"ordering", 1}, {"number", 1}})

	cursor, err := eventCollection.Find(ctx, d, &queryOptions)
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

func GetEventByMeetingAndNumber(id string, number int) (model.Event, error) {
	events, err := getEventsByBsonDocument(bson.D{{"meeting", id}, {"number", number}})
	if err != nil {
		return model.Event{}, err
	}

	if len(events) > 0 {
		return events[0], nil
	}

	return model.Event{}, errors.New("no entry with given meeting and number found")
}

func GetEventByMeetingAndNumberForLivetiming(id string, number int) (dto.EventLivetimingDto, error) {
	events, err := getEventsByBsonDocument(bson.D{{"meeting", id}})
	if err != nil {
		return dto.EventLivetimingDto{}, err
	}

	if len(events) <= 0 {
		return dto.EventLivetimingDto{}, errors.New("no entry with given meeting found")
	}

	var eventLivetiming dto.EventLivetimingDto

	for i := 0; i < len(events); i++ {
		if events[i].Number == number {
			eventLivetiming.Event = events[i]
			if i > 0 {
				eventLivetiming.PrevEvent = events[i-1]
			}
			if i < len(events)-1 {
				eventLivetiming.NextEvent = events[i+1]
			}
			break
		}
	}

	return eventLivetiming, nil
}

func GetEventsAsPartsByMeetId(id string) ([]model.MeetingPart, error) {
	events, err := getEventsByBsonDocument(bson.D{{"meeting", id}})
	if err != nil {
		return []model.MeetingPart{}, err
	}

	var partMap = make(map[int]model.MeetingPart)

	for _, event := range events {
		if &event.Part == nil {
			continue
		}
		_, has := partMap[event.Part.Number]
		if !has {
			event.Part.Events = []model.Event{}
			partMap[event.Part.Number] = event.Part
		}
		part := partMap[event.Part.Number]
		part.Events = append(part.Events, event)

		partMap[event.Part.Number] = part
	}

	var parts []model.MeetingPart

	for _, part := range partMap {
		parts = append(parts, part)
	}

	sort.SliceStable(parts, func(i, j int) bool {
		return parts[i].Number < parts[j].Number
	})

	return parts, nil
}

func GetEventsByMeetId(id string) ([]model.Event, error) {
	return getEventsByBsonDocument(bson.D{{"meeting", id}})
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

func ImportEvent(event model.Event, styleName string, PartNumber int) (*model.Event, bool, error) {
	// check if event with meeting and number exists
	// if not, create
	//		match style
	//		find part
	// 		set ordering
	// else exit
	existing, err := GetEventByMeetingAndNumber(event.Meeting, event.Number)
	if err != nil {
		if err.Error() == "no entry with given meeting and number found" {
			// create
			style, err2 := GetStyleByName(styleName)
			if err2 != nil {
				return nil, false, err2
			}
			event.Style = style

			event.Part = model.MeetingPart{
				Number: PartNumber,
			}

			event.Ordering = event.Number

			newEvent, err3 := AddEvent(event)
			if err3 != nil {
				return nil, false, err3
			}
			return &newEvent, true, nil
		}
		return nil, false, err
	}
	return &existing, false, nil
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
