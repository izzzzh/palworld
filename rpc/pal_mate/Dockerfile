FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
COPY . .
RUN go mod tidy
RUN go mod download
COPY ./rpc/pal_mate/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/pal_mate ./rpc/pal_mate


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/pal_mate /app/pal_mate
COPY --from=builder /app/etc /app/etc

CMD ["./pal_mate", "-f", "etc/palmate.yaml"]
