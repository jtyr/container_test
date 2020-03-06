from golang AS builder

ARG elfkickers_version=3.1

ADD http://www.muppetlabs.com/~breadbox/pub/software/ELFkickers-${elfkickers_version}.tar.gz /tmp
RUN tar -C /tmp -xf /tmp/ELFkickers-${elfkickers_version}.tar.gz
RUN make -C /tmp/ELFkickers-${elfkickers_version}/

COPY main.go /
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o /container_test /main.go
RUN /tmp/ELFkickers-${elfkickers_version}/sstrip/sstrip -z /container_test


FROM scratch
MAINTAINER Jiri Tyr

COPY --from=builder /container_test /

EXPOSE 8080

ENTRYPOINT ["/container_test"]
