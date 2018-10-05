FROM golang:1.9.7-alpine
ADD . /go/src/github.com/dsmontoya/axolog
WORKDIR /go/src/github.com/dsmontoya/axolog
RUN apk add --update git
RUN go get -v
RUN go get github.com/smartystreets/goconvey
RUN go install
CMD axolog
