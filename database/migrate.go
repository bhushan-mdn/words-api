package database

import (
	"context"
	"fmt"

	"github.com/bhushan-mdn/words-api/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func Migrate() {

	coll := utils.GetCollection()

	// Count all documents in the `words` collection
	count, err := coll.CountDocuments(context.TODO(), bson.D{})
	utils.CheckError(err)
	fmt.Println(count)

	// Check if data is already present
	if count == 0 {

		words := utils.GetWords()

		// Insert words into the coll
		for _, word := range words {
			_, err = coll.InsertOne(context.TODO(), word)
			utils.CheckError(err)
		}

	} else {
		fmt.Println("Data is already present in the collection.")
	}
}
