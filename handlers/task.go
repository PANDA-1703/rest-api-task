package handlers

import (
	"context"
	database "rest-api-go/db"
	"rest-api-go/models"
	"time"

	"github.com/gofiber/fiber/v3"
)

// POST - создание задачи
func CreateTask(c fiber.Ctx) error {
	var task models.Task
	if err := c.Bind().Body(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный запрос"})
	}

	query := `INSERT INTO tasks (title, description, status, created_at, updated_at) 
				VALUES ($1, $2, $3, NOW(), NOW())
				RETURNING id, created_at, updated_at`

	err := database.DB.QueryRow(context.Background(), query,
		task.Title, task.Description, task.Status,
	).Scan(&task.ID, &task.Created_at, &task.Updated_at)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка создания задачи", "details": err.Error()})
	}

	return c.Status(201).JSON(task)

}

// GET - получение всех задач
func GetTasks(c fiber.Ctx) error {
	rows, err := database.DB.Query(context.Background(), `SELECT id, title, description, status, created_at, updated_at FROM tasks`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка запроса к БД!", "details": err.Error()})
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Ошибка чтения данных!", "details": err.Error()})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

// GET - получение задачи по id
func GetTask(c fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task

	err := database.DB.QueryRow(context.Background(),
		`SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1`,
		id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Задача не найдена!", "details": err.Error()})
	}

	return c.JSON(task)
}

// PUT - обновление задачи по id
func PutTask(c fiber.Ctx) error {
	id := c.Params("id")
	var data models.Task

	if err := c.Bind().Body(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный запрос!", "details": err.Error()})
	}

	var updatedAt time.Time
	err := database.DB.QueryRow(context.Background(), `
		UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
		`, data.Title, data.Description, data.Status, id).Scan(&updatedAt)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Задача не найдена!", "details": err.Error()})
	}

	return c.Status(200).JSON("Задача обновлена")
}

// DELETE - удаление задачи по id
func DeleteTask(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec(context.Background(), `DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка удаления задачи!", "details": err.Error()})
	}

	return c.Status(204).JSON("Задача удалена")
}
