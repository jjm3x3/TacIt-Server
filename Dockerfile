FROM us.gcr.io/tacit-196502/go-build:master
LABEL maintainer="jjm3333@gmail.com"

COPY . /root/src/tacit-api

RUN cd /root/src/tacit-api && \
    dep ensure && \
    go build && \
    cp tacit-api /bin/

CMD ["/bin/tacit-api"]
