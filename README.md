# REST API для управления задачами (TODO-лист).

## Установка

1. Склонировать репозиторий:

```bash
git clone https://github.com/PANDA-1703/rest-api-task.git
cd rest-api-go
```

2. Установить зависимости:
```bash
go mod download
``` 

3. Создать `.env` и прописать 
```bash
# Порт сервера
PORT=3000

# Настройки БД 
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres    
DB_PASSWORD=your_password
DB_NAME=your_db
```

4. Применить миграцию:
```bash
migrate -source file://migrations -database "postgres://postgres:your_password@localhost/your_db?sslmode=disable" up
```

или создать таблицу вручную:
```SQL
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT CHECK(status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

5. Запустить сервер:
```bash
go run main.go
# или 
go run .
```

### Запуск через Docker
*В `docker-compose.yml` указать свои данные БД.*

```bash
docker compose up -d --build
```

API позволяет:
- Создавать задачу.
- Читать список задач.
- Читать одну конкретную задачу.
- Обновлять задачу.
- Удалять задачу.

## Функционал

| **Метод** | **URL**           | **Описание**                |
|-------|---------------|-------------------------|
| POST  | /tasks        | Создание задачи         |
| GET   | /tasks        | Получение всех задач    |
| GET   | /tasks/:id    | Получение задачи по id  |
| PUT   | /tasks/:id    | Обновление задачи       |
| DELETE| /tasks/:id    | Удаление задачи         |

## Стек

- **Go** 1.24
- **Fiber** v3.0.0-beta.4
- **PostgreSQL** 17
- **pgx** v5.7.4
- golang-migrate/migrate
- Docker/Docker-compose

## Структура БД

|Поле|ТИП|ОПИСАНИЕ|
|------|--------|------|
|id|SERIAL|Уникальный номер|
|title|TEXT NOT NULL|Название задачи|
|description|TEXT|Описание задачи|
|status|TEXT|Статус (new/in_progress/done)|
|created_at|TIMESTAMP|Дата создания|
|updated_at|TIMESTAMP|Дата обновления|


## Примеры запросов

- **POST /tasks - создать задачу**
```bash
curl -X POST http://localhost:3000/tasks -H "Content-type: application/json" -d '{"title":"Купить хлеб","description":"Сходить в магазин за хлебом","status":"new"}'
```

- **GET /tasks — получить все задачи**
```bash
curl http://localhost:3000/tasks
```

- **PUT /tasks/1 — обновить задачу**
```bash
curl -X PUT http://localhost:3000/tasks/1 -H "Content-Type: application/json" -d '{"title":"Купить хлеб и молоко","description":"Сходить в магазин, купить хлеб и молоко","status":"in_progress"}'
```

- **DELETE /tasks/1 — удалить задачу**
```bash
curl -X DELETE http://localhost:3000/tasks/1
```

## Структура проекта
```
rest-api-go
 ┣ db
 ┃ ┗ database.go
 ┣ handlers
 ┃ ┗ task.go
 ┣ migrations
 ┃ ┣ 001_create_tasks_table.down.sql
 ┃ ┗ 001_create_tasks_table.up.sql
 ┣ models
 ┃ ┗ task.go
 ┣ routes
 ┃ ┗ tasks.go
 ┣ .env
 ┣ .gitignore
 ┣ docker-compose.yml
 ┣ Dockerfile
 ┣ go.mod
 ┣ go.sum
 ┣ main.go
 ┗ README.md
```