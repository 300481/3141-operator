FROM golang:1.12.9-buster AS builder

ARG KUBECTL_VERSION=v1.15.2
ARG HELM_VERSION=v2.14.3

WORKDIR /usr/local/bin
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl
RUN chmod +x kubectl
RUN wget https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz
RUN tar xvzf helm-${HELM_VERSION}-linux-amd64.tar.gz
RUN mv linux-amd64/helm .

WORKDIR /go/src/github.com/300481/
RUN git clone https://github.com/300481/3141-operator.git
WORKDIR /go/src/github.com/300481/3141-operator/
RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o /usr/local/bin/3141-operator .

FROM alpine:3.10.1

WORKDIR /

COPY --from=builder /usr/local/bin/ /usr/local/bin/
COPY scripts /scripts

RUN apk add --no-cache \
        ca-certificates \
        bash

ENTRYPOINT [ "/usr/local/bin/3141-operator" ]