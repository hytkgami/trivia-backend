FROM golang:1.22.5-alpine

RUN go install github.com/air-verse/air@latest & go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air", "-c", ".air.toml"]
