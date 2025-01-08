FROM golang:1.23.1 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

# Build a static binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY ./config ./config

# Ensure binary is executable
RUN chmod +x ./main

EXPOSE 8089

CMD ["./main", "-config=./config/config.yml"]
