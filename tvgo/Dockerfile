FROM golang:1.24.4-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV PATH=$PATH:/go/bin

COPY . .

RUN swag init -g main.go -o docs

RUN go build -o tvgo ./main.go

RUN chmod +x start.sh

EXPOSE 8083

CMD ["./start.sh"]
