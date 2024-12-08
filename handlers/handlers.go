package handlers

import (
	"strconv"
	"todo/database"
	"todo/models"

	"github.com/gofiber/fiber/v3"
)

func GetTasks(c fiber.Ctx) error {
	tasks, err := database.GetTasks()
	if err != nil {
		return c.Status(500).SendString("Something went wrong, failed get tasks")
	}
	return c.JSON(tasks)
}

func GetTask(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(401).SendString("Invalid task id")
	}
	task, err := database.GetTaskByID(id)
	if err != nil {
		return c.Status(404).SendString("Task id not found")
	}
	return c.JSON(task)
}

func CreateTask(c fiber.Ctx) error {
	var task models.Task
	if err := c.Bind().Body(&task); err != nil {
		return c.Status(401).SendString("Invalid task data")
	}
	if err := database.Create(task); err != nil {
		return c.Status(500).SendString("Something went wrong, failed create task")
	}
	return c.SendStatus(201)
}

func DeleteTask(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(401).SendString("Invalid task id")
	}
	if err := database.Delete(id); err != nil {
		return c.Status(404).SendString("Task id not found, failed delete task")
	}
	return c.SendStatus(200)
}
