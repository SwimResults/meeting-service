package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Event struct {
	Identifier    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`                       // automatically
	Number        int                `json:"number,omitempty" bson:"number,omitempty"`                 // PDF + DSV
	Distance      int                `json:"distance,omitempty" bson:"distance,omitempty"`             // PDF + DSV
	RelayDistance string             `json:"relay_distance,omitempty" bson:"relay_distance,omitempty"` // PDF + DSV
	Meeting       string             `json:"meeting,omitempty" bson:"meeting,omitempty"`               // import service
	Gender        string             `json:"gender,omitempty" bson:"gender,omitempty"`                 // PDF + DSV
	StyleId       primitive.ObjectID `json:"-" bson:"style_id,omitempty"`                              // automatically
	Style         Style              `json:"style,omitempty" bson:"-"`                                 // PDF + DSV
	Final         EventFinal         `json:"final,omitempty" bson:"final,omitempty"`                   // manually
	Part          MeetingPart        `json:"part,omitempty" bson:"part,omitempty"`                     // PDF + DSV
	Finished      bool               `json:"finished,omitempty" bson:"finished,omitempty"`             // PDF + DSV + Livetiming
	Certified     bool               `json:"certified,omitempty" bson:"certified,omitempty"`           // PDF + DSV ! used for result file links
	Ordering      int                `json:"ordering,omitempty" bson:"ordering,omitempty"`             // automatically / manually
	AddedAt       time.Time          `json:"added_at,omitempty" bson:"added_at,omitempty"`             // automatically
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`         // automatically
}

type EventFinal struct {
	IsPrelim bool   `json:"is_prelim,omitempty" bson:"is_prelim,omitempty"`
	IsFinal  bool   `json:"is_final,omitempty" bson:"is_final,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
}

type MeetingPart struct {
	Number int     `json:"number,omitempty" bson:"number,omitempty"`
	Name   string  `json:"name,omitempty" bson:"name,omitempty"` // purpose?!
	Events []Event `json:"events,omitempty" bson:"-"`
}
