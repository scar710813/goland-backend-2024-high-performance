FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build ./cmd/server

FROM scratch
WORKDIR /app
CMD ["go", "run", "./cmd/server/main.go"]