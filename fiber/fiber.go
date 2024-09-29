package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "pong"})
    })

    app.Listen(":8080")
}
