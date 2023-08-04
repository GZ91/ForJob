FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY internal/* ./
COPY cmd/shortener/main.go ./

RUN go build -o /go-dockerapp

EXPOSE 8080

CMD [ "/go-dockerapp" ]