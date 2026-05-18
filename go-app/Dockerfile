FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o server .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server ./server

EXPOSE 8000

CMD ["./server"]
