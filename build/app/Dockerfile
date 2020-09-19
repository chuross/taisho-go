FROM golang:1.15.2

ARG NAME_SPACE

EXPOSE 8080

RUN go get github.com/oxequa/realize

RUN mkdir -p "/go/src/${NAME_SPACE}"
WORKDIR "/go/src/${NAME_SPACE}"

CMD ["realize", "start"]