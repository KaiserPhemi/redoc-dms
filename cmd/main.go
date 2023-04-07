package main

import (
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	Api := app.Group("/api/v1")
	app.Listen(":7777")
}
