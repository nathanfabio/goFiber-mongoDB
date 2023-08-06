package tags

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tags struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
	Tasks []string `json:"-" bson:"tasks,omitempty"`
}