package tags

import (
	"github.com/nathanfabio/goFiber-mongoDB/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findOrCreate(name string) (doc Tags, err error) {
	filter := bson.M{"name": name}

	err = db.FindOnlyOne("tags", filter, &doc)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}

	if doc.Name != "" {
		return
	}

	doc.Name = name

	id, err := db.Insert("tags", &doc)
	if err != nil {
		return
	}

	doc.ID = id

	return
}