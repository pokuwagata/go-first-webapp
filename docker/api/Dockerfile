FROM golang:1.12

ADD . /go/src/first_webapp
WORKDIR /go/src/first_webapp

RUN GO111MODULE=on go build server.go

CMD env=$ENV ./server $PORT
