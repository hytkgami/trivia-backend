FROM mcr.microsoft.com/devcontainers/go:0-1.20

USER vscode

RUN go install github.com/cosmtrek/air@latest & go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air", "-c", ".air.toml"]
