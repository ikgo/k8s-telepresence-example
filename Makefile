REGISTRY?=gcr.io/k8s-example

build: build-user build-widget

build-user:
	docker build ./user_service -t ${REGISTRY}/user-service

build-widget:
	docker build ./widget_service -t ${REGISTRY}/widget-service

run:
	kubectl create -f user-service.yaml -f widget-service.yaml

swap-user:
	go build -o user-service user_service/main.go
	telepresence \
	--swap-deployment user-service \
	--expose 8080 \
	--run ./user-service \
	--port=8444 \
	--method vpn-tcp

swap-widget:
	go build -o widget-service widget_service/main.go
	telepresence \
	--swap-deployment widget-service \
	--expose 8080 \
	--run ./widget-service \
	--port=8444 \
	--method vpn-tcp

restart-pods:
	kubectl delete pods --all

print-services:
	minikube service --url user-service
	minikube service --url widget-service

clean:
	rm -f user-service
	rm -f widget-service
	rm -f telepresence.log

.PHONY: build build-user build-widget