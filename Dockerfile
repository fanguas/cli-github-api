FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o cli-github-api .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cli-github-api .

CMD ["./cli-github-api"]