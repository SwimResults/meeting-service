package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Meeting struct {
	Identifier primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DateStart  time.Time          `json:"date_start,omitempty" bson:"date_start,omitempty"`
	DateEnd    time.Time          `json:"date_end,omitempty" bson:"date_end,omitempty"`
	LocationId primitive.ObjectID `json:"-" bson:"location_id,omitempty"`
	Location   Location           `json:"location,omitempty" bson:"-"`
	Organizer  primitive.ObjectID `json:"organizer_id,omitempty" bson:"organizer_id,omitempty"`
	SeriesId   primitive.ObjectID `json:"-" bson:"series_id,omitempty"`
	Series     MeetingSeries      `json:"series,omitempty" bson:"-"`
	Iteration  int                `json:"iteration,omitempty" bson:"iteration,omitempty"`
	State      string             `json:"state,omitempty" bson:"state,omitempty"` // options: HIDDEN; ANNOUNCED; PREPARATION; OPENING; RUNNING; BREAK; PAUSE; FINAL; OVER; ARCHIVED;
	MeetId     string             `json:"meet_id,omitempty" bson:"meet_id,omitempty"`
	Data       MeetingData        `json:"data,omitempty" bson:"data,omitempty"`
	Layout     MeetingLayout      `json:"layout,omitempty" bson:"layout,omitempty"`
	AddedAt    time.Time          `json:"added_at,omitempty" bson:"added_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
