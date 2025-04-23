# syntax=docker/dockerfile:1.4
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY ad/go.mod ad/go.sum ./
COPY proto ../proto
# Copy vendor folder if using local proto module
COPY ad/vendor/ ad/vendor/

# Force Go to use vendor folder
ENV GOFLAGS="-mod=vendor"

RUN go mod tidy
# Copy the source code
COPY ad/. .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
# Final stage: minimal image
FROM alpine:latest

WORKDIR /root/

# Copy built binary
COPY --from=builder /app/main .

# Expose gRPC port (change to 50052 if needed)
EXPOSE 8000
# Run the service
CMD ["./main"]
