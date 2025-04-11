# People API Service

Микросервис для управления данными людей с обогащением информации (возраст, пол, национальность) через внешние API.

## 📌 Возможности

- Создание, чтение, обновление и удаление записей о людях
- Фильтрация и пагинация списка людей
- Автоматическое обогащение данных: 
добавление age (api.agify.io), gender (api.genderize.io), nationality (api.nationalize.io)
- Документированный REST API (Swagger)

## 🚀 Быстрый старт
### Установка
```bash
# Клонировать репозиторий
git clone https://github.com/ninaprokopova/testTaskEffectiveMobile

# Скопировать .env.example
cp .env.example .env
```
### Заполнить настройки в .env
Ниже пример из файла .env.example
```
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=mypass
DB_NAME=habrdb
DB_SSLMODE=disable

# Server
SERVER_PORT=8080

# LogLevel
LOG_LEVEL=debug
```
### Запуск
```bash
# Установить зависимости
go mod download

# Запустить миграции и сервер
go run cmd/main.go
```
## 📌 Документация API
После запуска сервера документация доступна по адресу:
http://localhost:8080/swagger/index.html

## 📁 Структура проекта
```text
.
├── cmd/               # Точка входа
├── config/            # Конфигурация
├── docs/              # Swagger docs
├── internal/
│   ├── dto/           # Data Transfer Objects
│   ├── handlers/      # HTTP обработчики
│   └── service/       # Бизнес-логика
├── packages/
│   └── database/      # Модели и миграции
├── .env.example       # Шаблон конфига
├── go.mod
└── README.md
```