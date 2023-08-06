package tasks

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tasks struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Tags        []string           `json:"tags" bson:"tags,omitempty"`
	Assing     primitive.ObjectID `json:"assing" bson:"assing,omitempty"`
	Done        bool               `json:"done" bson:"done,omitempty"`
}