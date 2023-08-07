package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Find retrieves all documents from the specified collection in the database.
func Find(collection string, filter bson.M, documents any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	return cursor.All(context.Background(), documents)
}

//FindOne retrieves a document from the specified collection by its ID.
func FindOne(collection string, id string, document any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	return c.FindOne(context.Background(), filter).Decode(document)
}

//FindOnlyOne finds a single document in a collection based on the given filter and decodes it into the provided document object.
func FindOnlyOne(collection string, filter bson.M, document any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	return c.FindOne(context.Background(), filter).Decode(document)
}