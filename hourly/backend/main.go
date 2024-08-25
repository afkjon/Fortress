package main

import (
	"fmt"
	"log"
	"os"

	"github.com/afkjon/Fortress/hourly/db"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"

	"github.com/markbates/goth/providers/google"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8001/auth/google/callback"),
	)

	// Create DB Connection
	_, err := db.DbConnection(os.Getenv("DSN"))
	if err != nil {
		fmt.Println("DB Driver creation failed", err.Error())
		return
	}

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	// Setup Logger
	logFile, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	app.Use(logger.New(logger.Config{
		Output: logFile,
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
