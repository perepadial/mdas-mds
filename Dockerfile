FROM golang:rc-alpine3.15

EXPOSE 8090

RUN apk add libc6-compat

RUN mkdir /server

COPY app /server

WORKDIR /server

CMD ./app 