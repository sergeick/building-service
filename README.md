# Building Service API

Этот проект предоставляет RESTful API для управления зданиями с использованием Go, Gin и Swagger для документации API.

## Содержание

- [Предварительные требования](#предварительные-требования)
- [Установка](#установка)
- [Запуск проекта](#запуск-проекта)
- [Документация API](#документация-api)
- [Тестирование API](#тестирование-api)
- [Лицензия](#лицензия)

## Предварительные требования

- Go 1.18+
- PostgreSQL

## Установка

1. **Клонируйте репозиторий:**

   ```sh
   git clone https://github.com/yourusername/building-service.git
   cd building-service
   ```

2. **Установите зависимости:**

   ```sh
   go mod tidy
   ```

3. **Сгенерируйте документацию Swagger:**

   ```sh
   swag init -g main.go
   ```

4. **Настройте базу данных:**

   - Создайте базу данных и пользователя PostgreSQL.
   - Обновите файл `config/config.yaml` вашими учетными данными для базы данных.

   Пример `config/config.yaml`:

   ```yaml
   database:
     host: localhost
     port: 5432
     user: your_db_user
     password: your_db_password
     dbname: your_db_name
   ```

5. **Запустите миграции базы данных:**

   ```sh
   psql -U your_db_user -d your_db_name -f migrations/001_create_buildings_table.up.sql
   ```

## Запуск проекта

Запустите сервер:

```sh
go run main.go
```

Сервер будет запущен на [http://localhost:8080](http://localhost:8080).

## Документация API

Документация API доступна по адресу [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

## Тестирование API

Вы можете протестировать конечные точки API, используя cURL. Вот несколько примеров команд:

1. Создать новое здание:

   ```sh
   curl -X POST http://localhost:8080/api/v1/buildings -H "Content-Type: application/json" -d '{"name": "Building A", "city": "City X", "year": 2020, "floors": 5}'
   ```

2. Получить список зданий:

   ```sh
   curl -X GET "http://localhost:8080/api/v1/buildings?city=CityX&year=2020&floors=5"
   ```

3. Просмотреть документацию Swagger:

   ```sh
   curl -X GET http://localhost:8080/swagger/index.html
   ```

4. Просмотреть сгенерированную документацию OpenAPI:

   ```sh
   curl -X GET http://localhost:8080/swagger/doc.json
   ```

 
## Лицензия

MIT
