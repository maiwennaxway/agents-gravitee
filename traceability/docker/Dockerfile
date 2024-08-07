# Build image
# golang:1.21.6-alpine3.19 linux/amd64
FROM docker.io/golang@sha256:2523a6f68a0f515fe251aad40b18545155135ca6a5b2e61da8254df9153e3648 AS builder

ENV GOFLAGS "-mod=mod"
ENV GOWORK "off"

RUN apk --no-cache update && \
  apk --no-cache add -U bash build-base git

ENV BASEPATH=/opt/gravitee/agents-gravitee
ENV AGENT=${BASEPATH}/traceability

RUN mkdir -p ${BASEPATH}

WORKDIR ${BASEPATH}

# Copy necessary files
COPY . .

RUN ls

WORKDIR ${AGENT}

RUN export time=`date +%Y%m%d%H%M%S` && \
  export commit_id=`git rev-parse --short HEAD` && \
  export version=`git tag -l --sort='version:refname' | grep -Eo '[0-9]{1,}\.[0-9]{1,}\.[0-9]{1,3}$' | tail -1` && \
  export sdk_version=`go list -m github.com/Axway/agent-sdk | awk '{print $2}' | awk -F'-' '{print substr($1, 2)}'` && \
  export GOOS=linux && \
  export CGO_ENABLED=0 && \
  export GOARCH=amd64 && \
  go build -tags static_all \
  -ldflags="-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildTime=${time}' \
  -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildVersion=${version}' \
  -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildCommitSha=${commit_id}' \
  -X 'github.com/Axway/agent-sdk/pkg/cmd.SDKBuildVersion=${sdk_version}' \
  -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildAgentName=GraviteetraceabilityAgent'" \
  -a -o ${AGENT}/bin/gravitee_traceability_agent ${AGENT}/main.go

# Create non-root user
RUN addgroup -g 2500 axway && adduser -u 2500 -D -G axway axway
RUN chown -R axway:axway /opt/gravitee/agents-gravitee/traceability/bin/gravitee_traceability_agent
USER axway

# alpine 3.19 linux/amd64
FROM docker.io/alpine@sha256:13b7e62e8df80264dbb747995705a986aa530415763a6c58f84a3ca8af9a5bcd

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /opt/gravitee/agents-gravitee/traceability/bin/gravitee_traceability_agent /gravitee_traceability_agent

RUN mkdir /keys /specs && \
  chown -R axway /keys /specs && \
  apk --no-cache add openssl libssl3 libcrypto3 musl musl-utils libc6-compat busybox curl && \
  touch /gravitee_traceability_agent.yml && \
  find / -perm /6000 -type f -exec chmod a-s {} \; || true

USER axway

VOLUME ["/keys", "/specs"]

HEALTHCHECK --retries=1 CMD /gravitee_traceability_agent --status || exit 1

ENTRYPOINT ["/gravitee_traceability_agent"]