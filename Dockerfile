FROM golang:1.10.0-stretch
LABEL maintainer="jjm3333@gmail.com"

COPY TacIt ./src/bin

CMD ["./bin/TacIt"]
