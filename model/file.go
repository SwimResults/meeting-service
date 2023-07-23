package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type StorageFile struct {
	Identifier primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Meeting    string             `json:"meeting,omitempty" bson:"meeting,omitempty"`
	Path       string             `json:"path,omitempty" bson:"path,omitempty"`
	Url        string             `json:"url,omitempty" bson:"url,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Extension  string             `json:"extension,omitempty" bson:"extension,omitempty"`
	InFileList bool               `json:"in_file_list,omitempty" bson:"in_file_list,omitempty"`
	Hidden     bool               `json:"hidden,omitempty" bson:"hidden,omitempty"`
	Existing   bool               `json:"existing,omitempty" bson:"existing,omitempty"`
	Downloads  int                `json:"downloads,omitempty" bson:"downloads,omitempty"`
	AddedAt    time.Time          `json:"added_at,omitempty" bson:"added_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
