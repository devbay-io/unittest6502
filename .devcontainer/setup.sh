#!/usr/bin/env bash

sudo apt update
sudo apt install pre-commit 64tass -y

curl -sS https://starship.rs/install.sh -o /tmp/install.sh
chmod +x /tmp/install.sh
/tmp/install.sh -y
mkdir -p ~/.config/
cp -rp ./.devcontainer/starship.toml ~/.config/starship.toml

echo 'eval "$(starship init bash)"' >> ~/.bashrc

go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/go-critic/go-critic/cmd/gocritic@latest
go install github.com/BurntSushi/toml/cmd/tomlv@master
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.58.2
