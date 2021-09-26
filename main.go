package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO: Organize into modules: routes, database, controllers

func main() {
	rand.Seed(time.Now().UnixNano())

	var option string

	flag.StringVar(&option, "o", "serve", "Specify serve or migrate.")

	flag.Parse()

	if option == "migrate" {
		migrate()
	} else {
		fmt.Println("Using old data.")
	}

	coll := getCollection()

	var words []Word
	cursor, err := coll.Find(context.TODO(), bson.D{})
	CheckError(err)

	err = cursor.All(context.TODO(), &words)
	CheckError(err)

	fmt.Println("Server running at localhost" + SERVER_PORT)

	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New())

	// TODO: solidify the current routes, add more if necessary

	app.Get("/api/random", func(c *fiber.Ctx) error {
		word := randomWord(words)

		return c.JSON(word)
	})

	app.Get("/api/words", func(c *fiber.Ctx) error {
		query := c.Query("count", "1")
		count, err := strconv.Atoi(query)
		if err != nil {
			log.Fatal(err)
		}
		resp := randomWords(words, count)

		return c.JSON(resp)
	})

	app.Get("/api/available/words", func(c *fiber.Ctx) error {
		var resp []string
		for _, v := range words {
			resp = append(resp, v.Name)
		}

		return c.JSON(resp)
	})

	app.Get("/api/word/:word", func(c *fiber.Ctx) error {
		req := c.Params("word")
		fmt.Println(req)
		// TODO: Find a more elegant approach for this.
		var resp Word
		for i := range words {
			if words[i].Name == req {
				resp = words[i]
				break
			}
		}

		return c.JSON(resp)
	})

	log.Fatal(app.Listen(SERVER_PORT))
}
