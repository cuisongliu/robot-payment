FROM golang:1.12.7
COPY ../../../robot /go/src/github.com/fanux/
WORKDIR /go/src/github.com/fanux/robot/example/go-restful
RUN go build -mod vendor -o robot
CMD ./robot