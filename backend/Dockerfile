FROM golang:1.8.0-alpine

ENV APP_NAME go-docker
ENV REPO_BASE_URL github.com/whimatthew
ENV APP_PATH $GOPATH/src/$REPO_BASE_URL/$APP_NAME/backend

COPY . $APP_PATH
WORKDIR $APP_PATH
RUN apk add --update rsync && \
    rsync -a $APP_PATH/vendor/* $GOPATH/src/

CMD ["sh", "-c", "go run $(echo ${APP_PATH}/src/app.go)"]