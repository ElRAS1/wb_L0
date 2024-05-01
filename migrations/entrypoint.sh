
#!/bin/sh
# Загрузка переменных из файла.env
set -a
source /migrations/.env
set +a

# Использование переменных в команде migrate
migrate -path=/migrations -database="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up
