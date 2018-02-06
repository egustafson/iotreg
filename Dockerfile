FROM golang:alpine

WORKDIR /go/src/iotreg
copy . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 4480

CMD ["iotreg"]
