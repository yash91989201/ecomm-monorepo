# Builder stage 
FROM golang:1.23.0-alpine AS builder
WORKDIR /

# Download dependencies for the common module
COPY common ./common
WORKDIR /common
RUN go mod download

# Download dependencies for rest module
WORKDIR /
COPY . .
WORKDIR /gateways/rest
RUN go mod download

# Build the final binary for rest gateway
WORKDIR /
RUN go build -o rest-gateway ./gateways/rest/cmd/rest

# Final stage 
FROM alpine:3.21.0
WORKDIR /
COPY --from=builder /rest-gateway ./
EXPOSE 8080
CMD ["./rest-gateway"]
