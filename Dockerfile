FROM golang:1.9.2-alpine3.6
MAINTAINER Jevon Wild

ADD . /go/src/github.com/dorkusprime/og-gofer

RUN cd /go/src/github.com/dorkusprime/og-gofer && go install -v .

CMD /go/bin/og-gofer

EXPOSE 8080
