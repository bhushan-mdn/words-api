package routes

import (
	"github.com/bhushan-mdn/words-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetParamsWord(c *fiber.Ctx) error {

	req := c.Params("word")

	resp := utils.WordsMap[req]

	return c.JSON(resp)
}
