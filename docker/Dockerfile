# Dockerfile for Go application with "air" auto-reloading
FROM golang:latest

RUN pwd

WORKDIR /app

COPY .. .

RUN go mod download

# # Install "air" for auto-reloading
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]