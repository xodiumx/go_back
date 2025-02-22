FROM golang:1.23-alpine AS builder

WORKDIR /src

# Копируем go.mod и go.sum, затем загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарный файл
RUN go build -o main .

# Используем минимальный образ для финального контейнера
FROM alpine:latest

WORKDIR /root/

# Копируем скомпилированный бинарник из предыдущего образа
COPY --from=builder /src/main .

# Запускаем приложение
CMD ["./main"]
