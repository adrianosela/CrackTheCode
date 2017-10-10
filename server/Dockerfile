FROM alpine:3.5

RUN apk add --update bash curl && rm -rf /var/cache/apk/*

ADD server /bin/server

EXPOSE 8080

CMD ["/bin/server"]
