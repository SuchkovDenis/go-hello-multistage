FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o hello-go main.go

FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/hello-go .
CMD ["./hello-go"]
