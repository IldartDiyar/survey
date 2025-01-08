FROM golang:1.23.1 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY ./config ./config

EXPOSE 8089

CMD ["./main", "-config=./config/config.yml"]
