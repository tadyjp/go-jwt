FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /go/src/app
RUN go get \
    github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
