# Build image
# golang:1.21.3-alpine3.18 linux/amd64
FROM docker.io/golang@sha256:27c76dcf886c5024320f4fa8ceb57d907494a3bb3d477d0aa7ac8385acd871ea AS builder

# Installer Git dans l'image de base
RUN apk --no-cache update && \
    apk --no-cache add -U bash build-base git

# Définir les variables d'environnement
ENV BASEPATH=/go/src/github.com/maiwennaxway/agents-gravitee
ENV AGENT=${BASEPATH}/discovery

# Créer le répertoire de travail
RUN mkdir -p ${AGENT}
WORKDIR ${AGENT}

# Copier les fichiers nécessaires
COPY . .

# Initialiser Git et le module Go
RUN git init && \
    git remote add origin https://github.com/maiwennaxway/agents-gravitee && \
    git fetch --tags \
    go mod init gravitee_discovery_module

# Extraire les informations de Git
RUN time=$(date +%Y%m%d%H%M%S) && \
    commit_id=$(git rev-parse --short HEAD) && \
    version=$(git describe --tags --abbrev=0) && \
    sdk_version=$(go list -m github.com/Axway/agent-sdk | awk '{print $2}' | awk -F'-' '{print substr($1, 2)}') && \
    GOOS=linux && \
    CGO_ENABLED=0 && \
    GOARCH=amd64 && \
    go build -tags static_all \
    -ldflags="-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildTime=${time}' \
    -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildVersion=${version}' \
    -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildCommitSha=${commit_id}' \
    -X 'github.com/Axway/agent-sdk/pkg/cmd.SDKBuildVersion=${sdk_version}' \
    -X 'github.com/Axway/agent-sdk/pkg/cmd.BuildAgentName=graviteeDiscoveryAgent'" \
    -o bin/gravitee_discovery_agent main.go

# Créer un utilisateur non root
RUN addgroup -g 2500 axway && adduser -u 2500 -D -G axway axway
RUN chown -R axway:axway bin/gravitee_discovery_agent

# Alpine 3.18.3
FROM docker.io/alpine@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33

# Copier les fichiers de l'étape précédente
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/github.com/maiwennaxway/agents-gravitee/discovery/bin/gravitee_discovery_agent /gravitee_discovery_agent

# Créer des répertoires nécessaires et installer des dépendances
RUN mkdir /keys && \
    chown -R axway /keys && \
    apk --no-cache add openssl libssl3 libcrypto3 musl musl-utils libc6-compat busybox curl && \
    touch /gravitee_discovery_agent.yml && \
    find / -perm /6000 -type f -exec chmod a-s {} \; || true

# Utiliser l'utilisateur non root
USER axway

# Définir le volume
VOLUME ["/keys"]

# Healthcheck
HEALTHCHECK --retries=1 CMD /gravitee_discovery_agent --status || exit 1

# Point d'entrée
ENTRYPOINT ["/gravitee_discovery_agent"]
