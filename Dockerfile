FROM golang

ENV PORT "8080"
ADD ./towncenter /go/bin/towncenter
ADD ./config.json /go/bin/config.json

ENTRYPOINT /go/bin/towncenter

EXPOSE 8080