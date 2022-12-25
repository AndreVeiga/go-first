FROM golang:1.17

WORKDIR /go-fisrt

COPY go.mod ./
COPY hello.go ./

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]