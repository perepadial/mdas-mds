FROM golang:rc-alpine3.15

EXPOSE 8090

RUN mkdir /server

RUN apk add git

RUN git clone  https://github.com/perepadial/mdas-mds.git /server

WORKDIR /server

RUN go build -o app .

CMD ./app 