package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteOne deletes a document from the specified collection based on its ID.
func DeleteOne(collection, id string) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	result, err := c.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("%d documents deleted", result.DeletedCount)
	}
	return nil
}