FROM golang:1.14-alpine3.11 AS builder

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o calculator cmd/*.go

CMD ["./calculator"]