FROM golang:1.22.1-alpine

WORKDIR /app

RUN apk add --no-cache make git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make swag_gen

RUN go build -o tvgo ./main.go

RUN chmod +x start.sh

EXPOSE 8083

CMD ["./start.sh"]
