# Use official Golang image as the build environment
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o simx-go-todo ./cmd/api

# Use a minimal image for running
FROM alpine:latest
WORKDIR /root/

# Install ca-certificates for HTTPS, and PostgreSQL client for debugging
RUN apk --no-cache add ca-certificates postgresql-client

COPY --from=builder /app/simx-go-todo .

EXPOSE 8080

CMD ["./simx-go-todo"]
