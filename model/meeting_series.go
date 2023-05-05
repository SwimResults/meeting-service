package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MeetingSeries struct {
	Identifier primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NameFull   string             `json:"name_full,omitempty" bson:"name_full,omitempty"`
	NameMedium string             `json:"name_medium,omitempty" bson:"name_medium,omitempty"`
	NameShort  string             `json:"name_short,omitempty" bson:"name_short,omitempty"`
	LocationId primitive.ObjectID `json:"-" bson:"location_id,omitempty"`
	Location   Location           `json:"location,omitempty" bson:"-"`
	Organizer  primitive.ObjectID `json:"organizer_id,omitempty" bson:"organizer_id,omitempty"`
	AddedAt    time.Time          `json:"added_at,omitempty" bson:"added_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
