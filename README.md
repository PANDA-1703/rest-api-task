# REST API для управления задачами (TODO-лист).

## Установка

1. Клонируйте репозиторий:


## Эндпоинты

| Метод | URL           | Описание                |
|-------|---------------|-------------------------|
| POST  | /tasks        | Создание задачи         |
| GET   | /tasks        | Получение всех задач    |
| PUT   | /tasks/:id    | Обновление задачи       |
| DELETE| /tasks/:id    | Удаление задачи         |

## Пример запроса

```json
{
  "title": "Купить хлеб",
  "description": "Срочно купить хлеб к ужину",
  "status": "new"
}

postgreSQL 17
go version go1.24.3 windows/amd64
github.com/gofiber/fiber/v3 v3.0.0-beta.4
github.com/jackc/pgx/v5 v5.7.4