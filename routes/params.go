package routes

import (
	"fmt"

	"github.com/bhushan-mdn/words-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetParamsWord(c *fiber.Ctx) error {
	words := utils.GetWords()
	req := c.Params("word")
	fmt.Println(req)
	// TODO: Find a more elegant approach for this.
	var resp utils.Word
	for i := range words {
		if words[i].Name == req {
			resp = words[i]
			break
		}
	}

	return c.JSON(resp)
}
