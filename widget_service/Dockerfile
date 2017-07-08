FROM golang:1.8

EXPOSE 8080

RUN [ "mkdir", "-p", "/go/src/github.com/johnpryan/k8s-dev-example/widget_service"]
WORKDIR /go/src/github.com/johnpryan/k8s-dev-example/widget_service
ADD . .
RUN go get .
RUN go build -o main main.go
CMD [ "./main" ]
