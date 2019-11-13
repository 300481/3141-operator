# Building stage
FROM golang:1.13.4-buster AS builder

WORKDIR /go/src/github.com/300481/cda
COPY . .
WORKDIR /go/src/github.com/300481/cda/cmd/cda
RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /usr/local/bin/cda

ARG KUBECTL_VERSION=v1.16.2
ARG HELM_VERSION=v2.16.0

WORKDIR /usr/local/bin
RUN wget -q https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    chmod +x kubectl && \
    wget -q https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    tar xvzf helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    mv linux-amd64/helm . && \
    rm -rf helm-${HELM_VERSION}-linux-amd64.tar.gz linux-amd64

# Run stage
FROM alpine:3.10.3

WORKDIR /

COPY --from=builder /usr/local/bin/ /usr/local/bin/

RUN apk add --no-cache \
        ca-certificates \
        bash \
        git

CMD [ "server" ]

ENTRYPOINT [ "/usr/local/bin/cda" ]