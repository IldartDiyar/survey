FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY ./config ./config

EXPOSE 8089

CMD ["./main", "-config=./config/config.yml"]
