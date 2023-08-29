FROM golang:1.20-alpine

LABEL authors="GeorgiyZ"
WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./GoDockerapp ./cmd/shortener/main.go
EXPOSE 8080

ENTRYPOINT ["./GoDockerapp"]