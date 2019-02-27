FROM golang:1.11.2

WORKDIR /go/src/go-service1

ADD . /go/src/go-service1

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init -v

CMD ["go","run","main.go"]
