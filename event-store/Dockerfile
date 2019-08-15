FROM golang:1.10.2-alpine AS build

RUN apk update && apk add git

#ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
#RUN chmod +x /usr/bin/dep

WORKDIR /go/src/github.com/krijnrien/guilddigger/event_store

#COPY Gopkg.toml Gopkg.lock ./

COPY . .
#RUN dep ensure --vendor-only
RUN go install ./...

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p usr/bin/configs/cert/


WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY --from=build /go/src/github.com/krijnrien/guilddigger/event-store/configs/cert configs/cert/

#CMD ["event-store -port=50051"]
CMD ["event-store"]