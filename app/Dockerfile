FROM golang:1.8.0-alpine

RUN apk add --update bash && \
    mkdir -p /app
COPY . /app
WORKDIR /app

RUN mv /app/go/vendor/* $GOPATH/src

CMD ["go", "run", "/app/go/app.go"]