# Install goose in builder stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application - собираем весь пакет cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd

# Final stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

# Copy goose from builder
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Copy binary and migrations
COPY --from=builder /app/bin/api .
COPY --from=builder /app/internal/adapters/postgres/migrations ./internal/adapters/postgres/migrations

EXPOSE 8080

CMD ["./api"]