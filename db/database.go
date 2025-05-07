package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

// Подключение к БД
func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка: Не удалось загрузить данные из .env!")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var connectErr error
	db, connectErr := pgx.Connect(context.Background(),
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", dbUser, dbPassword, dbName, dbHost, dbPort))
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v!\n", connectErr)
	}
	DB = db
	fmt.Println("Подключение к БД успешно.")
}
