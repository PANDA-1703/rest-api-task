package routes

import (
	"rest-api-go/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterTasksRoutes(app *fiber.App) {
	app.Get("/tasks", handlers.GetTasks)
	app.Get("/tasks/:id", handlers.GetTask)
	app.Post("/tasks", handlers.CreateTask)
	app.Put("/tasks/:id", handlers.PutTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

}
