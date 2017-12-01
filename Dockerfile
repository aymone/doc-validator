FROM golang:1.9.2

RUN mkdir -p /go/src/document-validator

COPY ./src/ /go/src/document-validator

WORKDIR /go/src/document-validator

RUN go get github.com/codegangsta/gin &&\
    go-wrapper download &&\
    go-wrapper install

# Expose gin proxy port
EXPOSE 3000

CMD gin run
