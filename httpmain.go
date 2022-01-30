package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/L"
	"log"
)

func main() {
	app := fiber.New()

	app.Post("/write", func(c *fiber.Ctx) error {
		var any interface{}
		if err := c.BodyParser(&any); err != nil {
			return err
		}
		L.Describe(any)
		return nil
	})

	log.Fatal(app.Listen(":8081"))
}
