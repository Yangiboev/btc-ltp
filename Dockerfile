FROM golang:1.22.3-bullseye AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0  GOOS=linux go build -o main main.go

FROM alpine:3.17.3

WORKDIR /app

COPY --from=builder /app/main .

ENV PORT=8080
EXPOSE 8080

CMD ["./main"]
