FROM golang as build

ENV GOPROXY=https://goproxy.cn

ADD . /base

WORKDIR /base

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o base ./cmd/base

FROM alpine:latest

RUN echo "https://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /base/base /usr/bin/base

RUN chmod +x /usr/bin/base

ENTRYPOINT ["base"]