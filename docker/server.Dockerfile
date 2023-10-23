FROM build:latest as BUILDER

FROM alpine:latest

RUN apk add openssl
# certs and executable in the same dir
COPY --from=BUILDER /build/bin/server* /usr/local/bin/


EXPOSE 4062

CMD [ "server" ]
