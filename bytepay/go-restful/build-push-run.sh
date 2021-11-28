rm robot
docker run --rm -v /Users/fanux/work/src/github.com/fanux/robot:/go/src/github.com/fanux/robot \
  -w /go/src/github.com/fanux/robot/example/go-restful golang:1.13.5  go build -o robot -mod vendor
docker build -t registry.cn-hangzhou.aliyuncs.com/sealyun/robot:latest .
docker push registry.cn-hangzhou.aliyuncs.com/sealyun/robot:latest
kubectl delete -f deploy.yaml || true
kubectl apply -f deploy.yaml