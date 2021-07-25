package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	fiberApp := fiber.New() // 'fiber app' is an instance of Fiber

	// Here we create a route using the Get HTTP request method
	fiberApp.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Welcome to my Fiber Web app")
	})

	fiberApp.Get("/info", func(c *fiber.Ctx) error {
		return c.SendString("Fiber is an Express inspired web framework built on top of Fasthttp," +
			"the fastest HTTP engine for Go.")
	})

	fiberApp.Get("/fiber", func(c *fiber.Ctx) error {
		return c.Redirect("https://gofiber.io/", 301)
	})

	fiberApp.Get("/video", func(c *fiber.Ctx) error {
		return c.Redirect("https://www.youtube." +
			"com/watch?v=kvwsPeWDLM8&t=209s", 301)
	})

	// Here we assign a port number to our webapp
	log.Fatal(fiberApp.Listen(":2021"))
}

// Our HTTP request method is the 'GET'
// 'path' is a virtual path on the server
/* 'func(*fiber.Ctx) error' is a callback function containing the Context
executed when the route is matched
*/

