ARG NAME_SPACE="github.com/chuross/taisho"
ARG PROJECT_DIR="/go/src/${NAME_SPACE}"

###

FROM golang:1.15.2 as builder

ARG NAME_SPACE
ARG PROJECT_DIR

RUN mkdir -p ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}

COPY ./ ${PROJECT_DIR}

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o taisho ./cmd/taisho/main.go

###

FROM alpine:latest

ARG PROJECT_DIR

RUN mkdir -p /go/bin
WORKDIR /go/bin

COPY --from=builder ${PROJECT_DIR}/taisho ./taisho

EXPOSE 8080

CMD ["./taisho"]
