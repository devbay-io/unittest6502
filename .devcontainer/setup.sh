#!/usr/bin/env bash

sudo apt update
sudo apt install pre-commit -y

go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/go-critic/go-critic/cmd/gocritic@latest
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.58.2
