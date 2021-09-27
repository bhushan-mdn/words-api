package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGODB_STRING string = "mongodb://localhost:27017"
var MONGODB_DATABASE string = "testdb"
var MONGODB_COLLECTION string = "words"
var SERVER_PORT string = "5000"

type Word struct {
	Name    string `json:"word"`
	Meaning string `json:"meaning"`
}

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var Words []Word
var WordsList []string
var WordsMap = make(map[string]Word)

func GetCollection() *mongo.Collection {
	fmt.Println(MONGODB_STRING)
	// Set client options
	clientOptions := options.Client().ApplyURI(MONGODB_STRING)

	// Connect to MongoDB
	client, e := mongo.Connect(context.TODO(), clientOptions)
	CheckError(e)

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	CheckError(e)

	// Get collection as reference
	return client.Database(MONGODB_DATABASE).Collection(MONGODB_COLLECTION)
}

func GetWords() []Word {
	coll := GetCollection()

	var words []Word
	cursor, err := coll.Find(context.TODO(), bson.D{})
	CheckError(err)

	err = cursor.All(context.TODO(), &words)
	CheckError(err)

	return words
}

func GetWordsList() []string {
	var wordsList []string

	for _, v := range Words {
		wordsList = append(wordsList, v.Name)
	}

	return wordsList
}

func GetWordsMap() map[string]Word {

	for _, v := range Words {
		WordsMap[v.Name] = v
	}

	return WordsMap
}

func ParseWords() []Word {
	file, err := os.ReadFile("assets/words.json")
	CheckError(err)

	var words []Word
	err = json.Unmarshal(file, &words)
	CheckError(err)

	return words
}

func RandomWord(data []Word) Word {
	idx := rand.Intn(len(data))
	return data[idx]
}

func RandomWords(data []Word, count int) []Word {
	var words []Word

	for i := 0; i < count; i++ {
		idx := rand.Intn(len(data))
		words = append(words, data[idx])
	}

	return words
}
