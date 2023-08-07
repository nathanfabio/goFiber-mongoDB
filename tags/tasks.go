package tags

import (
	"sort"

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

func AddTask(taskID string, names []string) error {
	for _, name := range names {
		tag, err := findOrCreate(name)
		if err != nil {
			return err
		}

		i := sort.SearchStrings(tag.Tasks, taskID)

		if i < len(tag.Tasks) && tag.Tasks[i] == taskID {
			continue	
		}

		tag.Tasks = append(tag.Tasks, taskID)

		sort.Strings(tag.Tasks)

		result := new(Tag)

		err = db.UpdateOne("tags", tag.ID.Hex(), tag, result)
		if err != nil {
			return err
		}
	}

	return nil
}