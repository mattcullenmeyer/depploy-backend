FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux go build -o bin/main ./cmd/local

EXPOSE 8080

CMD ["./bin/main"]
