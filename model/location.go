package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Identifier primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Street     string             `json:"street,omitempty" bson:"street,omitempty"`
	Number     string             `json:"number,omitempty" bson:"number,omitempty"`
	City       string             `json:"city,omitempty" bson:"city,omitempty"`
	PostalCode string             `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
}