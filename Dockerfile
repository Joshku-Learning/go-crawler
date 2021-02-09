FROM golang:1.15-alpine

# go get 會用到
RUN apk add git \
    && apk add --no-cache git
    && go get github.com/pilu/fresh \
    && apk add ca-certificates \
