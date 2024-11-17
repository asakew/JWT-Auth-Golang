FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o hello-world

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/hello-world .

CMD ["./hello-world"]
