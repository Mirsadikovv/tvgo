FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tvgo ./main.go

RUN swag init -g main.go -o ./docs

RUN chmod +x start.sh

EXPOSE 8083

CMD ["./start.sh"]

