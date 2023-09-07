# Dockerfile for Go application with "air" auto-reloading
FROM golang:latest

RUN pwd

WORKDIR /app

COPY .. .

RUN go mod download

RUN go install github.com/cespare/reflex@latest

CMD ["./tmp/cli", "migrate"]
CMD ["reflex", "-c", "reflex.conf"]