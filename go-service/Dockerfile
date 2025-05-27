FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gin-api .
RUN chmod +x gin-api

EXPOSE 9001
CMD ["./gin-api"]
