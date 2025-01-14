FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go_telepresence
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/go_telepresence .
EXPOSE 8080
CMD ["./go_telepresence"]