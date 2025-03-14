FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/api ./cmd/api
COPY internal ./internal

ENV CGO_ENABLED=0  

WORKDIR /app/cmd/api
RUN go build -ldflags="-s -w" -o /app/main .

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .

EXPOSE 8083

CMD ["./main"]
