package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Event struct {
	Identifier    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Number        int                `json:"number,omitempty" bson:"number,omitempty"`
	Distance      int                `json:"distance,omitempty" bson:"distance,omitempty"`
	RelayDistance string             `json:"relay_distance,omitempty" bson:"relay_distance,omitempty"`
	Meeting       string             `json:"meeting,omitempty" bson:"meeting,omitempty"`
	Gender        string             `json:"gender,omitempty" bson:"gender,omitempty"`
	StyleId       primitive.ObjectID `json:"-" bson:"style_id,omitempty"`
	Style         Style              `json:"style,omitempty" bson:"-"`
	Final         EventFinal         `json:"final,omitempty" bson:"final,omitempty"`
	Part          MeetingPart        `json:"part,omitempty" bson:"part,omitempty"`
	Ordering      int                `json:"ordering,omitempty" bson:"ordering,omitempty"`
	AddedAt       time.Time          `json:"added_at,omitempty" bson:"added_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type EventFinal struct {
	IsPrelim bool   `json:"is_prelim,omitempty" bson:"is_prelim,omitempty"`
	IsFinal  bool   `json:"is_final,omitempty" bson:"is_final,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
}

type MeetingPart struct {
	Number int     `json:"number,omitempty" bson:"number,omitempty"`
	Name   string  `json:"name,omitempty" bson:"name,omitempty"`
	Events []Event `json:"events,omitempty" bson:"-"`
}
