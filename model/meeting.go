package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Meeting struct {
	Identifier primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	DateStart  primitive.Timestamp `json:"date_start,omitempty" bson:"date_start,omitempty"`
	DateEnd    primitive.Timestamp `json:"date_end,omitempty" bson:"date_end,omitempty"`
	LocationId primitive.ObjectID  `json:"-" bson:"location_id,omitempty"`
	Location   Location            `json:"location,omitempty" bson:"-"`
	Organizer  primitive.ObjectID  `json:"organizer_id,omitempty" bson:"organizer_id,omitempty"`
	SeriesId   primitive.ObjectID  `json:"-" bson:"series_id,omitempty"`
	Series     MeetingSeries       `json:"series,omitempty" bson:"-"`
	Iteration  int                 `json:"iteration,omitempty" bson:"iteration,omitempty"`
	State      string              `json:"state,omitempty" bson:"state,omitempty"`
}
