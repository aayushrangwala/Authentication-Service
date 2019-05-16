FROM golang:latest as builder
WORKDIR /go/src/github.com/aayushrangwala/Authentication-Service
ADD ./ .
RUN rm -rf ./vendor
RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN cd cmd/server && go build -o main

FROM alpine
WORKDIR /app 
COPY --from=builder /go/src/github.com/aayushrangwala/Authentication-Service/cmd/server/main .
RUN chmod ugo+rwx ./main
ENTRYPOINT ./main