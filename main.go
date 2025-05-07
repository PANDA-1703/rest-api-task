package main

import (
	"fmt"
	"log"
	"os"
	"rest-api-go/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"

	database "rest-api-go/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка: Не удалось загрузить данные из .env!")
	}

	dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	database.Connect()

	runMigrations(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName))

	app := fiber.New()

	routes.RegisterTasksRoutes(app)

	errApp := app.Listen(fmt.Sprintf(":%v", port))
	if errApp != nil {
		panic(errApp)
	}

}

func runMigrations(connectStr string) {
	m, err := migrate.New(
		"file://migrations",
		connectStr,
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v!\n", err)
	}

	if err := m.Up(); err != nil {
		if err.Error() != "no change" {
			log.Fatalf("Ошибка выполнения миграции: %v!\n", err)
		}
		fmt.Println("Нет изменений.")
	} else {
		fmt.Println("Миграция прошла успешно")
	}
}
