package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Incident struct {
	Identifier primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Meeting    string             `json:"meeting,omitempty" bson:"meeting,omitempty"`
	Type       string             `json:"type,omitempty" bson:"type,omitempty"` // EVENT, DURATION
	Name       string             `json:"name,omitempty" bson:"name,omitempty"` // Einschwimmen, Kampfrichtersitzung, Mannschaftsleistersitzung
	Start      time.Time          `json:"start,omitempty" bson:"start,omitempty"`
	End        time.Time          `json:"end,omitempty" bson:"end,omitempty"`
	PrevEvent  int                `json:"prev_event,omitempty" bson:"prev_event,omitempty"`
	NextEvent  int                `json:"next_event,omitempty" bson:"next_event,omitempty"`
	AddedAt    time.Time          `json:"added_at,omitempty" bson:"added_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
