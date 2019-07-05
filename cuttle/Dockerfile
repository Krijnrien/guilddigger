FROM golang:1.10.2-alpine AS build

RUN apk update && apk add git

#ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
#RUN chmod +x /usr/bin/dep

WORKDIR /go/src/github.com/krijnrien/cuttle

#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure --vendor-only
COPY . .
RUN go install ./...


FROM alpine
WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY --from=build /go/src/github.com/krijnrien/cuttle/configs configs


RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

#ADD /configs/cert/myCA.crt /usr/local/share/ca-certificates/myCA.crt
#RUN chmod 644 /usr/local/share/ca-certificates/myCA.crt && update-ca-certificates

ENV LOGLEVEL="DEBUG"
CMD ["cuttle"]