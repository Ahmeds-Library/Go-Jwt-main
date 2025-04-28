FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-jwt-app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY .env .

COPY --from=builder /app/go-jwt-app .

EXPOSE 8002

CMD ["./go-jwt-app"]
