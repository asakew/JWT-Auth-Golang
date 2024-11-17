# Dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o hello-world

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/hello-world .

CMD ["./hello-world"]
