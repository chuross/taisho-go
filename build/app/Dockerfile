FROM golang:1.15.2

ARG NAME_SPACE

RUN go get github.com/oxequa/realize

RUN mkdir -p "/go/src/${NAME_SPACE}"
WORKDIR "/go/src/${NAME_SPACE}"

CMD ["realize", "start"]