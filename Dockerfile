# ==============================
# Этап сборки
# ==============================
FROM golang:1.25-alpine AS builder

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник
RUN go build -o main ./cmd

# ==============================
# Финальный образ
# ==============================
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk --no-cache add ca-certificates tzdata

# Рабочая директория
WORKDIR /app

# Копируем бинарник и папку cmd
COPY --from=builder /app/main .
COPY --from=builder /app/cmd ./cmd

# Открываем порты, если нужно
EXPOSE 3003 3004

# Запуск приложения с флагом
CMD ["./main", "--use-local-env"]