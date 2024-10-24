FROM golang:1.23.2
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
COPY cmd cmd
COPY internal internal
EXPOSE 8080

RUN go build -o main cmd/app/main.go

CMD ["./main"]