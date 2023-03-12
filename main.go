package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/shahriarsohan/go-blog-practise/database"
	"github.com/shahriarsohan/go-blog-practise/routes"
)

func main() {
	database.Connect()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.SetUp(app)

	app.Listen(":" + port)

}
