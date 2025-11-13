FROM golang:1.25.2-bookworm

ARG APP_ROOT

ENV TZ=Asia/Tokyo \
    LANG=ja_JP.UTF-8 \
    LANGUAGE=ja_JP:ja \
    LC_ALL=ja_JP.UTF-8

# 共通パッケージをまとめてインストール
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      locales \
      vim \
      ripgrep \
      sudo \
      curl \
      ca-certificates \
      gnupg \
      python3 \
      iptables \
      ipset \
      iproute2 \
      dnsutils \
      jq \
      aggregate && \
    localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8 && \
    curl -sL "https://github.com/go-task/task/releases/download/v3.9.0/task_linux_amd64.deb" -o task.deb && \
    dpkg -i task.deb && \
    rm task.deb && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Node.js (v22) インストール
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - && \
    apt-get update && \
    apt-get install -y nodejs && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Codexを公式ガイド通りにグローバルインストール
RUN npm i -g @openai/codex

# Go ツール
RUN go install github.com/air-verse/air@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/rinchsan/gosimports/cmd/gosimports@latest

# 開発ユーザー作成
RUN groupadd -g 1000 dev && \
    useradd -m -u 1000 -g dev -s /bin/bash dev && \
    chown -R dev:dev /home/dev && \
    chown -R dev:dev /go/pkg/

# VS Code Server 用ディレクトリ
# VS Code Server 用のディレクトリを事前に作成して権限を設定
RUN mkdir -p /home/dev/.vscode-server/bin /home/dev/.vscode-server/extensions && \
    chown -R dev:dev /home/dev/.vscode-server

# Copy and set up firewall script
COPY .devcontainer/infra/go/init-firewall.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/init-firewall.sh && \
  echo "dev ALL=(root) NOPASSWD: /usr/local/bin/init-firewall.sh" > /etc/sudoers.d/go-firewall && \
  chmod 0440 /etc/sudoers.d/go-firewall

WORKDIR $APP_ROOT

USER dev

# bashrc 設定
COPY --chown=dev:dev .devcontainer/infra/go/.bashrc /home/dev/.bashrc

# Claude Code インストール
RUN curl -fsSL https://claude.ai/install.sh | bash && \
    echo 'export PATH="$HOME/.local/bin:$PATH"' >> /home/dev/.bashrc
