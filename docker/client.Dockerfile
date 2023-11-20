FROM build:latest as BUILDER

FROM alpine:latest

RUN apk add openssl
    # apk add --no-cache bash

COPY --from=BUILDER /build/bin/client /usr/local/bin/
COPY --from=BUILDER /build/bin/ddosclient /usr/local/bin/

ENTRYPOINT [ "tail", "-f", "/dev/null" ]
