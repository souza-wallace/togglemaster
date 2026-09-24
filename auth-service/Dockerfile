# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git ca-certificates

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /auth-service .

# Runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates wget

WORKDIR /app
COPY --from=builder /auth-service /app/auth-service

EXPOSE 8001

CMD ["/app/auth-service"]
