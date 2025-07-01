FROM golang:1.24.4-alpine

# Устанавливаем git, так как он нужен для go install
RUN apk add --no-cache git

WORKDIR /app

# Копируем go.mod и go.sum и скаливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копирум весь исходный код
COPY . .

# Устанавливаем swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Убедимся, что $GOPATH/bin добавлен в $PATH
ENV PATH=$PATH:/go/bin

# Генерируем Swagger-документацию
RUN swag init -g main.go -o ./docs

# Компилируем приложение
RUN go build -o tvgo ./main.go

# Делаем start.sh исполняемым
RUN chmod +x start.sh


# Открываем порт
EXPOSE 8083

# Запускаем приложение
CMD ["./start.sh"]