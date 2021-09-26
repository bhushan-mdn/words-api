package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func migrate() {

	coll := getCollection()

	// Count all documents in the `words` collection
	count, err := coll.CountDocuments(context.TODO(), bson.D{})
	CheckError(err)
	fmt.Println(count)

	// Check if data is already present
	if count == 0 {

		file := getWords()

		words := parseWords(file)

		// Insert words into the coll
		for _, word := range words {
			_, err = coll.InsertOne(context.TODO(), word)
			CheckError(err)
		}

	} else {
		fmt.Println("Data is already present in the collection.")
	}
}
