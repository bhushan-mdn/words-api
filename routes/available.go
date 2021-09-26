package routes

import (
	"github.com/bhushan-mdn/words-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAvailableWords(c *fiber.Ctx) error {
	words := utils.GetWords()
	var resp []string
	for _, v := range words {
		resp = append(resp, v.Name)
	}

	return c.JSON(resp)
}
