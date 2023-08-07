FROM golang:1.20-alpine

LABEL authors="GeorgiyZ"
RUN apk add git

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
COPY . .
RUN go mod download

#RUN go get -t ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o ./GoDockerapp ./cmd/shortener/main.go

EXPOSE 8080

ENTRYPOINT ["./GoDockerapp"]
#CMD ["./go-dockerapp"]