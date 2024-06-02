#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "Set up tooling dependencies globally"

PROTOC_VERSION=${PROTOC_VERSION:-27.0}

architecture=""
case $(uname -m) in
  i386 | i686)           architecture="x86_32" ;;
  x86_64)                architecture="x86_64" ;;
  arm | arm64 | aarch64) architecture="aarch_64" ;;
esac

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  protoc_bin=linux-$architecture
elif [[ "$OSTYPE" == "darwin"* ]]; then
  protoc_bin=osx-$architecture
else
  protoc_bin=linux-$architecture
fi

curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-$protoc_bin.zip"
unzip "protoc-$PROTOC_VERSION-$protoc_bin.zip" -d "$HOME/.local"
rm "protoc-$PROTOC_VERSION-$protoc_bin.zip"

echo ""
echo -e '\e[46mMake sure your shell startup script contains following lines, adding to PATH to execute the commands:\e[0m'
echo ""
# shellcheck disable=SC2016
echo -e '\t\e[36mexport PATH="$PATH:$HOME/.local/bin"\e[0m'
# shellcheck disable=SC2016
echo -e '\t\e[36mexport PATH="$PATH:$(go env GOPATH)/bin"\e[0m'
echo ""

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0

echo "Installed all dependencies"
echo ""
echo "Validate version v$PROTOC_VERSION by running"
(set -o xtrace; protoc --version)
echo ""
