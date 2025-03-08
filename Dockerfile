FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/api ./cmd/api
COPY internal ./internal

ENV CGO_ENABLED=0  

WORKDIR /app/cmd/api
RUN go build -ldflags="-s -w" -o /app/main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE    

CMD ["sh", "-c", "./main"]
