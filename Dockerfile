FROM golang:1.20

LABEL authors="GeorgiyZ"
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV BASE_URL=http://localhost:8080/
ENV DSN_Address=sbg3.sminex.com:5432
ENV DSN_BaseName=reductor_db
ENV DSN_Login=reductor_user
ENV DSN_Password=xcvM5KDXqEXX
ENV DSN_Sslmode=disable
ENV LOG_LEVEL=error
ENV Root_Token=Root_Token
ENV SERVER_ADDRESS=localhost:8080

#RUN go get -t ./...

#RUN CGO_ENABLED=0 GOOS=linux go build -o ./GoDockerapp ./cmd/shortener/main.go

RUN go build -o ./GoDockerapp ./cmd/shortener/main.go

EXPOSE 8080

#ENTRYPOINT ["./GoDockerapp"]
CMD ["./GoDockerapp"]