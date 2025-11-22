# E-commerce API Study

Мой первый API проект на Golang для изучения разработки backend-приложений.

## Стек технологий

- **Go** - основной язык программирования
- **PostgreSQL** - база данных
- **Docker** - контейнеризация
- **Docker Compose** - оркестрация контейнеров

## Запуск проекта

### Локальный запуск

1.1. Установите зависимости:

```bash
go mod download
```

1.2. Создайте `.env` файл для локального запуска:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=ecom
```

1.3. Запустите приложение:

```bash
go run main.go
```

### Запуск в Docker

2.1. Создайте `.env` файл для Docker:

```env
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=ecom
```

2.2. Запустите контейнеры:

```bash
docker compose up -d
```

> **Важно:** Для локального и Docker-запуска требуются разные конфигурации `.env` файлов, особенно параметр `DB_HOST`.
