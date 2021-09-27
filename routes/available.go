package routes

import (
	"github.com/bhushan-mdn/words-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAvailableWords(c *fiber.Ctx) error {

	return c.JSON(utils.WordsList)
}
