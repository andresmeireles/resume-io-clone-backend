# Dockerfile for Go application with "air" auto-reloading
FROM golang:latest

RUN pwd

WORKDIR /app

COPY .. .

RUN go mod download

# Install "air" for auto-reloading
RUN go install github.com/cosmtrek/air@latest
# Install delve debuger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["go build", "-o", "tmp/cli", "cmd/cli"]

CMD ["air"]