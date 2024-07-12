FROM alpine:3.15

LABEL maintainer="jonsosnyan <https://jonssonyan.com>"

WORKDIR /h-ui

ENV TZ=Asia/Shanghai
ENV GIN_MODE=release

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

COPY build/h-ui-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} h-ui

RUN apk update && apk add --no-cache bash tzdata ca-certificates nftables \
    && rm -rf /var/cache/apk/* \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && chmod +x /h-ui/h-ui

CMD ["./h-ui"]