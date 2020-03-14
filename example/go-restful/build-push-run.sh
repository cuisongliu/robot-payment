docker run --rm -v /Users/fanux/work/src/github.com/fanux/robot:/go/src/github.com/fanux/robot \
  -w /go/src/github.com/fanux/robot/example/go-restful golang:1.12.7  go build -o robot
docker build -t fanux/robot .
docker push fanux/robot
kubectl apply -f deploy.yaml