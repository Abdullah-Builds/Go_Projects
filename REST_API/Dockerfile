# Build stage
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .
COPY config ./config

EXPOSE 8082

CMD ["./server", "-config", "config/local.yaml"]