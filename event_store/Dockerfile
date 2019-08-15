FROM golang:1.10.2-alpine AS build

RUN apk update && apk add git

#ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
#RUN chmod +x /usr/bin/dep

WORKDIR /go/src/github.com/krijnrien/microguild

#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure --vendor-only
COPY . .
RUN go install ./...

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p usr/bin/configs/cert/


WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY --from=build /go/src/github.com/krijnrien/microguild/configs/cert configs/cert/
#COPY --from=build /go/src/github.com/krijnrien/microguild/configs/cert/myCA.crt /usr/local/share/ca-certificates/myCA.crt

#ADD /usr/bin/configs/cert/myCA.crt /usr/local/share/ca-certificates/myCA.crt
#RUN chmod 644 /usr/local/share/ca-certificates/myCA.crt && update-ca-certificates

RUN export "http_proxy=http://192.168.99.100:3128" \
    && export "https_proxy=https://192.168.99.100:3128" \
    && export "no_proxy=http://192.168.56.1" \