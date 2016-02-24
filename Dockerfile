FROM golang:1.6.0-alpine

RUN apk update && apk add curl git wget && rm -rf /var/cache/apk/*

RUN mkdir -p /app
ENV GOPATH /app
WORKDIR /app

EXPOSE 8080

CMD ["/app/run.sh"]

COPY . /app

