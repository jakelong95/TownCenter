FROM golang

ADD ./TownCenter /go/bin/towncenter
ADD ./config-dev.json /go/bin/config.json

ENTRYPOINT /go/bin/towncenter

EXPOSE 8084
