package routes

import (
	"log"
	"strconv"

	"github.com/bhushan-mdn/words-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetRandomWords(c *fiber.Ctx) error {
	words := utils.Words
	query := c.Query("count", "")
	if query == "" {
		return c.SendString("Provide count.")
	} else {
		count, err := strconv.Atoi(query)
		if err != nil {
			log.Fatal(err)
		}
		resp := utils.RandomWords(words, count)

		return c.JSON(resp)
	}
}
