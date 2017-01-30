FROM golang:1.8-alpine
ADD . /go/src/github.com/johnmcdnl/selenium
RUN go install github.com/johnmcdnl/selenium
ENTRYPOINT /go/bin/selenium