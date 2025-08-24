# -------- Build --------
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go binary
# CGO_ENABLED=0 disables CGO for static linking (good for Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# Build the Go tests binary
RUN CGO_ENABLED=0 GOOS=linux go test -c -o packer.test ./test

# -------- Run --------
FROM alpine:3.20

# Install certificates for HTTPS
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binaries from builder stage
COPY --from=builder /app/server .
COPY --from=builder /app/packer.test .

# Copy templates & static files
COPY web/ ./web/

# Expose port (matches your Go server ListenAndServe :8080)
EXPOSE 8080

# Run the server
CMD ["./server"]
