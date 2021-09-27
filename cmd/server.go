package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bhushan-mdn/words-api/routes"
	"github.com/bhushan-mdn/words-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

var Port string

// serveCmd represents the serve command
var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	// Long: `Starts the server`,
	Run: func(cmd *cobra.Command, args []string) {
		Server()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringVarP(&Port, "port", "p", "5000", "Enter the port for the server")

	utils.Words = utils.GetWords()

	utils.WordsList = utils.GetWordsList()

	utils.WordsMap = utils.GetWordsMap()
}

// TODO: Organize into modules: routes, database, controllers

func Server() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Server running at localhost" + ":" + Port)

	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New())

	// TODO: solidify the current routes, add more if necessary

	app.Get("/api/random", routes.GetRandomWord)

	app.Get("/api/words", routes.GetRandomWords)

	app.Get("/api/available/words", routes.GetAvailableWords)

	app.Get("/api/word/:word", routes.GetParamsWord)

	log.Fatal(app.Listen(":" + Port))
}
