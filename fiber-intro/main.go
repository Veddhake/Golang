package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error { // Simple get request with value hello world
	// 	return c.SendString("Hello, World!")
	// })

	// app.Get("/:value", func(c *fiber.Ctx) error { // returns the value we put inside of a url
	// 	return c.SendString("value: " + c.Params("value"))
	// })

	// app.Get("/:name?", func(c *fiber.Ctx) error { // question mark at the url endpoint is important so that even if no imput is given, still will execue where is john else will show hello name
	// 	if c.Params("name") != "" {
	// 		return c.SendString("Hello " + c.Params("name"))
	// 	}

	// 	return c.SendString("Where is John?")
	// })

	app.Static("/", "./public")

	app.Listen(":3000")
}
