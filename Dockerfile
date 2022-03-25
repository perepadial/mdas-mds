FROM golang:rc-alpine3.15

EXPOSE 8090

RUN mkdir /server

COPY ./app /server

WORKDIR /server

CMD ./app 