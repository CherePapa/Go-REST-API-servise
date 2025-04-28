package main

import (
	"log"
	"os"
	"web-service/database"
	"web-service/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	if err := godotenv.Load(); err != nil {
		log.Fatal("ERROR LOADING .end files")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
