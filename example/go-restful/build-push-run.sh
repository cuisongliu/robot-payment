docker build -t fanux/robot .
docker push fanux/robot
kubectl apply -f deploy.yaml