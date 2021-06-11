FROM golang:1.12

ADD . /go/src/github.com/igorok-follow/grpc-server

RUN go install github.com/igorok-follow/grpc-server@latest

ENTRYPOINT ["/go/bin/grpc-server"]

EXPOSE 65000