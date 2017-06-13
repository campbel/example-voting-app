FROM golang:1.8.3-alpine3.6
RUN apk update && apk add git && go get github.com/graarh/golang-socketio
ADD system_test.go /go/src/system/system_test.go
CMD go test system
