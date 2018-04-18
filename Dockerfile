FROM golang:1.10.0-stretch
LABEL maintainer="jjm3333@gmail.com"

COPY tacit-api ./src/bin

CMD ["./bin/tacit-api"]
