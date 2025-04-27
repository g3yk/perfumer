package handler

import (
	"API/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ResponseHTTP struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func GetEnv(c *fiber.Ctx) error {
	env := c.Params("env")
	res := config.Config(env)
	return c.SendString(fmt.Sprintf("Hello World: %s", res))
}
