FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tvgo ./main.go

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g main.go -o ./tvgo/docs

RUN chmod +x start.sh

EXPOSE 8083

CMD ["./start.sh"]

