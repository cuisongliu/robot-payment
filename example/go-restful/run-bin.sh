rm robot
docker run --rm -v /Users/fanux/work/src/github.com/fanux/robot:/go/src/github.com/fanux/robot \
  -w /go/src/github.com/fanux/robot/example/go-restful golang:1.12.7  go build -o robot
scp robot root@store.lameleg.com:/root