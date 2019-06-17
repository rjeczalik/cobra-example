FROM golang:1.12.6-alpine as builder
WORKDIR /go/example
COPY . /go/example
ENV GOBIN=/go/example/bin
RUN apk add --no-cache git gcc musl-dev linux-headers && go install -ldflags "-w -s" github.com/rjeczalik/cobra-example/cmd/example

FROM gruebel/upx:latest as upx
COPY --from=builder /go/example/bin/example /usr/bin/example
RUN upx --best --lzma /usr/bin/example

FROM alpine:3.9.4
LABEL authors "rafal@rjk.io"
COPY --from=upx /usr/bin/example /usr/bin/example
ENTRYPOINT []
CMD example
