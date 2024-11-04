FROM golang:1.22.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tvgo ./main.go

RUN chmod +x start.sh

EXPOSE 8083

CMD ["./start.sh"]
