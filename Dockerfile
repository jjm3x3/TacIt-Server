FROM alpine:latest
LABEL maintainer="jjm3333@gmail.com"

COPY tacit-api /bin/tacit-api

CMD ["/bin/tacit-api"]
