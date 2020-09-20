FROM golang:1.15.2 as builder

ARG NAME_SPACE

EXPOSE 8080

ARG NAME_SPACE="github.com/chuross/taisho"
ARG PROJECT_DIR="/go/src/${NAME_SPACE}"

RUN mkdir -p ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}

COPY ./ ${PROJECT_DIR}

RUN go build -o taisho ./cmd/taisho/main.go

###

FROM alpine:latest

RUN mkdir /workspace
WORKDIR /workspace

COPY --from=builder ${PROJECT_DIR}/taisho .

CMD ["./taisho"]
