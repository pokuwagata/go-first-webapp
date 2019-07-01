FROM golang

Add . /go/src/first_webapp
WORKDIR /go/src/first_webapp

RUN go install first_webapp

ENTRYPOINT /go/bin/first_webapp

EXPOSE 8080