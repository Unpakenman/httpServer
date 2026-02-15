# Используем официальный образ Go
FROM golang:1.25-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем приложение
RUN go build -o main ./cmd

# Финальный образ
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Копируем бинарный файл из builder
COPY --from=builder /app/main .

# Копируем папку cmd если там есть дополнительные файлы
COPY --from=builder /app/cmd ./cmd

# Указываем порт (если ваше приложение его использует)
EXPOSE 8080

# Команда для запуска (сохраняем ваш флаг --use-local-env)
CMD ["./main", "--use-local-env"]