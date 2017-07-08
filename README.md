
Start a minikube cluster:

```
minikube start
```

build the docker files in the minikube environment:

```
eval $(minikube docker-env)
make build
```

Start the services:

```
kubectl create -f user-service.yaml -f widget-service.yaml 
```