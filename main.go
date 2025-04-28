package main

import (
	"log"
	"os"
	"web-service/database"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	if err := godotenv.Load(); err != nil {
		log.Fatal("ERROR LOADING .end files")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)
}
