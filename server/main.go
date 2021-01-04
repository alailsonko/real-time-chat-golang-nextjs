package main

import (
	"fmt"
	"log"
	"os"

	"chat-server.com/auth"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Post("/signup", auth.SignUp)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

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
