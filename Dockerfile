
# Prepare binary
FROM golang:latest AS builder

WORKDIR /opt/workspace

RUN mkdir -p dns-ep/bin && mkdir -p dns-ep/pkg && mkdir -p dns-ep/src

COPY . dns-ep/

ENV GOPATH /opt/workspace/dns-ep
ENV GOBIN /opt/workspace/dns-ep/bin

WORKDIR /opt/workspace/dns-ep/src/dns-ep

RUN CGO_ENABLED=0 go install dns-ep.go

#########
# To obtain a small image, copy just binary into alpine image and use it
FROM alpine

COPY --from=builder /opt/workspace/dns-ep/bin/dns-ep /dns-ep

ENV DNSPORT "8080"

ENTRYPOINT /dns-ep
