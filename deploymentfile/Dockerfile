# Start from a minimal Go base image
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy source code
COPY . .

# Build the binary
RUN go build -o health-checker ./cmd

# Final image
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/health-checker .

# Run the binary
ENTRYPOINT ["./health-checker"]