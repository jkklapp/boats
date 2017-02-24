
FROM golang:1.8

RUN mkdir -p /app
COPY . /app
WORKDIR /app/go

RUN go get gopkg.in/gin-gonic/gin.v1 && \
    cd /app/go/ && \
    make build

CMD ["/app/go/bin/app.a"]