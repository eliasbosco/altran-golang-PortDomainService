FROM golang:1.15 as build-env
RUN apt-get update
RUN apt-get install python-dev libsasl2-dev sqlite3 -y

WORKDIR /go/src/app
ADD . /go/src/app

ARG GO111MODULE=off
ARG CGO_ENABLED=1

RUN GOPATH=/go/bin/app go get -d -v ./...
RUN GOPATH=/go/bin/app go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build-env /go/bin/app /
CMD ["/app"]
