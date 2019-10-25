FROM golang:latest

WORKDIR /go/src/app

COPY /go-rest-api/main.go .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]