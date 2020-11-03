#(C) Copyright 2020 Hewlett Packard Enterprise Development LP
FROM 657273346644.dkr.ecr.us-east-1.amazonaws.com/hpe-hcss/go-service-template-lib:v0.5.0-docs as docs
FROM 657273346644.dkr.ecr.us-west-2.amazonaws.com/mirror/alpine:3.12.0 as curl
RUN apk add --no-cache --update curl

FROM curl as grpc_health_probe
ENV GRPC_HEALTH_PROBE_VERSION v0.2.2
RUN curl -fsSL https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 -o /grpc_health_probe \
  && chmod +x /grpc_health_probe

FROM curl as tini
ENV TINI_VERSION v0.18.0
RUN curl -sSL https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static -o /sbin/tini && \
  chmod +x /sbin/tini

FROM 657273346644.dkr.ecr.us-east-1.amazonaws.com/hpe-hcss/go-build:v0.1.3 as builder
ARG GITHUB_TOKEN
ENV workspace /src
ENV GO111MODULE on
COPY go.* Makefile $workspace/
WORKDIR $workspace
RUN git config --global url.https://$GITHUB_TOKEN@github.com/.insteadOf https://github.com/ && make vendor
COPY . $workspace
RUN make all || exit 1

FROM alpine:3.12.0 as ca-certs
RUN apk add --no-cache --update ca-certificates

FROM scratch
COPY --from=builder /src/bin/hpcaas-job-scheduler /
COPY --from=grpc_health_probe /grpc_health_probe /usr/bin/
COPY --from=tini /sbin/tini /sbin/tini
COPY --from=docs /docs /docs
COPY --from=ca-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./etc/ /etc
USER hpcaas-job-scheduler
HEALTHCHECK NONE
ENTRYPOINT ["/sbin/tini","--","/hpcaas-job-scheduler"]
ARG TAG
ARG GIT_SHA
ARG GIT_DESCRIBE
ARG BUILD_DATE
ARG SRC_REPO
LABEL TAG=$TAG \
  GIT_SHA=$GIT_SHA \
  GIT_DESCRIBE=$GIT_DESCRIBE \
  BUILD_DATE=$BUILD_DATE \
  SRC_REPO=$SRC_REPO
