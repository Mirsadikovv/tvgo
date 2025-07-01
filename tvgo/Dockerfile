FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tvgo ./main.go

RUN chmod +x start.sh

EXPOSE 8083

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g main.go -o ./docs

CMD ["./start.sh"]

