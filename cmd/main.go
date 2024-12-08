package main

import (
	"fmt"
	"log"
	"os"
	"todo/database"
	"todo/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	database.ConnectDB(
		fmt.Sprintf(
			"host=database user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
		),
	)
	database.InitDB()

	app := fiber.New()
	app.Use(compress.New())
	app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/tasks", handlers.GetTasks)
	app.Get("/task/:id", handlers.GetTask)
	app.Post("/task", handlers.CreateTask)
	app.Delete("/task/:id", handlers.DeleteTask)

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}
}
