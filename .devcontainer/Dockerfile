# syntax=docker/dockerfile:1

FROM public.ecr.aws/docker/library/golang:1.24.3-bookworm

# Add binaries to PATH
ENV PATH="/root/.tenv/bin:/root/.local/bin:$(go env GOPATH)/bin:${PATH}"

# Install remaining development tools
RUN ARCH=$(dpkg --print-architecture) && \
    VERSION_CODENAME=$(. /etc/os-release && echo "$VERSION_CODENAME") && \
    DEBIAN_FRONTEND=noninteractive apt-get -qq update < /dev/null > /dev/null && apt-get -qq install ca-certificates curl < /dev/null > /dev/null && \
    echo 'deb [trusted=yes] https://apt.fury.io/caarlos0/ /' > /etc/apt/sources.list.d/caarlos0.list && \
    echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' > /etc/apt/sources.list.d/goreleaser.list && \
    install -m 0755 -d /etc/apt/keyrings && \
    curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc && \
    chmod a+r /etc/apt/keyrings/docker.asc && \
    echo "deb [arch=${ARCH} signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian ${VERSION_CODENAME} stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
    apt-get -qq update < /dev/null > /dev/null && apt-get -qq install \
    docker-ce-cli \
    docker-compose-plugin \
    git \
    gh \
    goreleaser \
    jq \
    pipx \
    svu \
    unzip \
    wget \
    zsh < /dev/null > /dev/null && \
    rm -rf /var/lib/apt/lists/* && \
    pipx install pre-commit < /dev/null > /dev/null && \
    pipx install go-task-bin < /dev/null > /dev/null

# Install Go tools
RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then ARCH="amd64"; elif [ "$ARCH" = "aarch64" ]; then ARCH="arm64"; fi && \
    # golangci-lint
    curl -sSL -o golangci-lint.tar.gz https://github.com/golangci/golangci-lint/releases/download/v2.1.2/golangci-lint-2.1.2-linux-${ARCH}.tar.gz && \
    tar --strip-components=1 -xzf golangci-lint.tar.gz -C /usr/local/bin golangci-lint-2.1.2-linux-${ARCH}/golangci-lint && \
    chmod +x /usr/local/bin/golangci-lint && \
    rm -fr golangci-lint.tar.gz && \
    # git-chglog
    curl -sSL -o git-chglog.tar.gz https://github.com/git-chglog/git-chglog/releases/download/v0.15.4/git-chglog_0.15.4_linux_${ARCH}.tar.gz && \
    tar xzf git-chglog.tar.gz -C /usr/local/bin git-chglog && \
    chmod +x /usr/local/bin/git-chglog && \
    rm git-chglog.tar.gz && \
    # goimports
    go install golang.org/x/tools/cmd/goimports@v0.31.0 > /dev/null && \
    # cosign
    LATEST_VERSION=$(curl https://api.github.com/repos/sigstore/cosign/releases/latest | grep tag_name | cut -d : -f2 | tr -d "v\", ") && \
    curl -O -L "https://github.com/sigstore/cosign/releases/latest/download/cosign_${LATEST_VERSION}_amd64.deb" && \
    dpkg -i cosign_${LATEST_VERSION}_amd64.deb && \
    rm -fr "comsign_${LATEST_VERSION}_amd64.deb" && \
    # tenv
    LATEST_VERSION=$(curl --silent https://api.github.com/repos/tofuutils/tenv/releases/latest | jq -r .tag_name) && \
    curl -O -L "https://github.com/tofuutils/tenv/releases/latest/download/tenv_${LATEST_VERSION}_amd64.deb" && \
    dpkg -i "tenv_${LATEST_VERSION}_amd64.deb" && \
    rm -fr /var/lib/apt/lists/* && \
    rm -fr "tenv_${LATEST_VERSION}_amd64.deb" && \
    # tofu
    tenv tofu install && \
    # terraform-docs
    curl -sSLo ./terraform-docs.tar.gz https://terraform-docs.io/dl/v0.20.0/terraform-docs-v0.20.0-$(uname)-amd64.tar.gz && \
    tar -xzf terraform-docs.tar.gz && \
    chmod +x terraform-docs && \
    mv terraform-docs /usr/local/bin/terraform-docs

# Set working directory
WORKDIR /workspace
