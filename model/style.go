package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Style struct {
	Identifier primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`       // like: BACKSTROKE, BUTTERFLY;
	Aliases    []string           `json:"aliases,omitempty" bson:"aliases,omitempty"` // different languages
	Relay      bool               `json:"relay,omitempty" bson:"relay,omitempty"`
}
