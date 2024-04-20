FROM golang:1.20 as builder

WORKDIR /app
COPY . .

RUN go build -trimpath -ldflags -o h-ui -ldflags="-s -w"

FROM alpine:3.15

RUN apk add --no-cache bash tzdata ca-certificates
ENV TZ=Asia/Shanghai
ENV GIN_MODE=release

WORKDIR /app
COPY --from=builder /app/h-ui .

ENTRYPOINT ["./h-ui"]