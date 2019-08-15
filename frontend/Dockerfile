FROM frolvlad/alpine-glibc

RUN mkdir -p /etc/envoy

ADD build_release_stripped/envoy /usr/local/bin/envoy\
COPY /go/src/github.com/krijnrien/guilddigger/frontend/docker/envoy/envoy.yaml /etc/envoy/envoy.yaml
#ADD configs/google_com_proxy.v2.yaml /etc/envoy/envoy.yaml

EXPOSE 10000

COPY ci/docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]