FROM golang:1.20

LABEL authors="GeorgiyZ"
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

#RUN go get -t ./...

#RUN CGO_ENABLED=0 GOOS=linux go build -o ./GoDockerapp ./cmd/shortener/main.go

RUN go build -o ./GoDockerapp ./cmd/shortener/main.go

EXPOSE 8080

#ENTRYPOINT ["./GoDockerapp"]
CMD ["./GoDockerapp"]