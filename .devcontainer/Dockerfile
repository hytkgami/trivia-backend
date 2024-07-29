FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

USER vscode

RUN go install github.com/air-verse/air@latest & go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air", "-c", ".air.toml"]
