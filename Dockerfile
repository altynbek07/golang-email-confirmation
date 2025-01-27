FROM golang:1.23.3-alpine

WORKDIR /app

COPY . .

CMD ["go", "run", "cmd/main.go"]