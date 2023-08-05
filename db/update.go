package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//UpdateOne updates a document in the specified collection based on the provided ID.
func UpdateOne(collection, id string, data, result any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	opts := options.FindOneAndUpdate().SetUpsert(false)

	err = c.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": data}, opts).Err()
	if err != nil {
		return err
	}

	return c.FindOne(context.Background(), filter).Decode(result)
}