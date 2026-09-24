# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /evaluation-service .

# Runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates wget

WORKDIR /app
COPY --from=builder /evaluation-service /app/evaluation-service

EXPOSE 8004

CMD ["/app/evaluation-service"]
