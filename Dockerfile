FROM alpine:latest
LABEL maintainer="jjm3333@gmail.com"

RUN apk update && apk add ca-certificates
COPY tacit-api /bin/tacit-api

CMD ["/bin/tacit-api"]
