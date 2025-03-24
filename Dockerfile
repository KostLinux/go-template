# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Install goose for migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/main.go

# Final stage
FROM golang:1.24-alpine

RUN apk add --no-cache ca-certificates tzdata && \
    mkdir -p /app/migrations

# Copy binaries and configurations
COPY --from=builder /app/api /app/
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY config.yaml /app/

WORKDIR /app

RUN chmod +x /app/api && \
    chmod +x /usr/local/bin/goose

EXPOSE 8080

CMD ["/app/api"]