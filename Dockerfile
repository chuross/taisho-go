FROM golang:1.15.2

ARG NAME_SPACE

EXPOSE 8080

ARG NAME_SPACE="github.com/chuross/taisho"
ARG PROJECT_DIR="/go/src/${NAME_SPACE}"

RUN mkdir -p ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}

COPY ./ ${PROJECT_DIR}

RUN go build -o taisho ./cmd/taisho/main.go

CMD ["bin/bash", "taisho"]