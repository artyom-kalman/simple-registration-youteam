# Простая регистрация с трёмя контейнерами

## Описание контейнеров

1. **Frontend (Простая форма регистрации)**

2. **Backend (Golang)**

3. **PostgreSQL**

## Как запустить проект

1. Клонировать репозиторий.
2. Создать .env файл
Пример.env:

  ```
  # Database
  POSTGRES_HOST=db
  POSTGRES_PORT=5432
  POSTGRES_DB=youteam
  POSTGRES_USER=postgres
  POSTGRES_PASSWORD=postgres

  # Go Server
  PORT=3030
  ```

4. Запустить проект командой:
   ```bash
   docker-compose up
   ```
