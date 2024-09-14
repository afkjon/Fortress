package main

import (
	"fmt"
	"log"
	"os"

	"github.com/afkjon/Fortress/hourly/backend/db"
	"github.com/afkjon/Fortress/hourly/backend/handlers"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	/* OAuth2.0
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8001/auth/google/callback"),
	)
	*/

	// Create DB Connection
	_, err := db.Connect(os.Getenv("DSN"))
	if err != nil {
		fmt.Println("DB Driver creation failed", err.Error())
		return
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("CLIENT_URI")},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{` +
			`"time":"${time_rfc3339_nano}",` +
			`"id":"${id}",` +
			`"remote_ip":"${remote_ip}",` +
			`"host":"${host}",` +
			`"method":"${method}",` +
			`"uri":"${uri}",` +
			`"user_agent":"${user_agent}",` +
			`"status":${status},` +
			`"error":"${error}",` +
			`"latency":${latency},` +
			`"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	// Setup Routes
	handlers.SetupRoutes(e)

	// Setup Logger
	logFile, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
