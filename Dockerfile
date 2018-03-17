FROM golang:alpine

ENV GOPATH /go 
ENV PATH $PATH:/go/bin

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get github.com/markbates/refresh

ADD . /go/src/github.com/weebagency/go-api-v2
RUN cd /go/src/github.com/weebagency/go-api-v2 && \
    go get ./...