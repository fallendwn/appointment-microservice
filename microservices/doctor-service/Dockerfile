FROM golang:1.25.1-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app
ENV GOPROXY=direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 50051
CMD ["./main"]