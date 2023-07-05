FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o gokedex

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gokedex .

EXPOSE 8080

CMD ["./gokedex"]