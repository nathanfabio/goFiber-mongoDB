package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Insert inserts data into the specified collection.
func Insert(collection string, data any) (primitive.ObjectID, error) {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	result, err := c.InsertOne(context.Background(), data)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}