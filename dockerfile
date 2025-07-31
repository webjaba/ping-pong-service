FROM golang:1.24.5-alpine AS builder
WORKDIR /app/
RUN pwd
COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

FROM scratch
WORKDIR /app/
COPY --from=builder /app/main .
CMD ["./main"]