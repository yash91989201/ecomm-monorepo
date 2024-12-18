# Builder stage 
FROM golang:1.23.0-alpine AS builder
WORKDIR /

# Download dependencies for the common module
COPY common ./common
WORKDIR /common
RUN go mod download

# Download dependencies for inventory module
WORKDIR /
COPY . .
WORKDIR /services/inventory
RUN go mod download

# Build the final binary for inventory service
WORKDIR /
RUN go build -o inventory_service ./services/inventory/cmd/inventory

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /inventory_service ./
EXPOSE 8080
CMD ["./inventory_service"]

