FROM golang

Add . /go/src/first_webapp
WORKDIR /go/src/first_webapp

RUN go build server.go

CMD ./server $PORT
