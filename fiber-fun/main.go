package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// ini adalah entry point program go
	app := fiber.New()

	// instance baru dari fiber melalui variable app
	app.Get("/", func(c *fiber.Ctx) error {
		// mengembalikan nilai string test saat mengakses route /
		return c.SendString("test")
	})

	// :name adalah sesuatu yang dikirimkan melalui sebuah variable di url (params)
	app.Get("/greet/:name", func(c *fiber.Ctx) error {
		// c.Params adalah function dari fiber untuk mengambil params yang ada di dalam url yang sudah kita tuliskan
		name := c.Params("name")
		return c.SendString("Hallo " + name)
	})

	log.Println("Server running on port 8080")
	app.Listen(":8080")
}
