package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":" + getPort()))
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		portEnv := "3030"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
		return portEnv
	}
	portEnv := port
	return portEnv
}
