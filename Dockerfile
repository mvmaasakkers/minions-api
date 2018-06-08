FROM alpine
WORKDIR /app

RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
COPY ./build/api /app/api

EXPOSE 80

ENTRYPOINT ["/app/api"]