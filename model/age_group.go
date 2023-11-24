package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AgeGroup struct {
	Identifier primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Meeting    string             `json:"meeting,omitempty" bson:"meeting,omitempty"`
	Event      int                `json:"event,omitempty" bson:"event,omitempty"`     // 15
	Default    bool               `json:"default,omitempty" bson:"default,omitempty"` // false
	MinAge     string             `json:"min_age,omitempty" bson:"min_age,omitempty"` // 2004
	MaxAge     string             `json:"max_age,omitempty" bson:"max_age,omitempty"` // 2002
	Ages       []string           `json:"ages,omitempty" bson:"ages,omitempty"`       // 2002, 2003, 2004
	IsYear     bool               `json:"is_year,omitempty" bson:"is_year,omitempty"` // true
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`       // Jahrgänge 2002 - 2004
	AddedAt    time.Time          `json:"added_at,omitempty" bson:"added_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
