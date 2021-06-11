FROM golang:1.16

ADD . /go/src/github.com/igorok-follow/grpc-server

RUN go install github.com/igorok-follow/grpc-server

ENTRYPOINT ["/go/bin/grpc-server"]

EXPOSE 5300