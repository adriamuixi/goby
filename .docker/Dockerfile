FROM golang:1.14
LABEL maintainer="Adrià Muixí <adriamuixi@gmail.com>"

ENV GO111MODULE=on

WORKDIR /go/src/github.com/goby/

COPY ./.docker/dev-entrypoint.sh /dev-entrypoint.sh
RUN chmod +x /dev-entrypoint.sh

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src /go/src/github.com/goby
RUN go build

CMD ["./goby"]
