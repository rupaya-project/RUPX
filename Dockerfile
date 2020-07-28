FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /rupaya
RUN cd /rupaya && make rupaya

FROM alpine:latest

WORKDIR /rupaya

COPY --from=builder /rupaya/build/bin/rupaya /usr/local/bin/rupaya

RUN chmod +x /usr/local/bin/rupaya

EXPOSE 7050
EXPOSE 9050

ENTRYPOINT ["/usr/local/bin/rupaya"]

CMD ["--help"]
