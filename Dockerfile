FROM golang:1.15-alpine

# go get 會用到
RUN apk add git \
    && go get github.com/pilu/fresh \
    && apk add ca-certificates \
