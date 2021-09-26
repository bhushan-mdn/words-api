package routes

import (
	"github.com/bhushan-mdn/words-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetRandomWord(c *fiber.Ctx) error {
	words := utils.GetWords()
	word := utils.RandomWord(words)

	return c.JSON(word)
}
