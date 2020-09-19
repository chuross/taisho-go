FROM golang:1.15.2

ARG NAME_SPACE

EXPOSE 8080

ARG PROJECT_DIR="/go/src/${NAME_SPACE}"

RUN mkdir -p ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}

COPY ./ .

RUN go build -o taisho ./cmd/taisho/main.go

CMD ["taisho"]
