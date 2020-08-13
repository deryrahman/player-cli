#!/bin/bash

OS=("android" "darwin" "darwin" "darwin" "darwin" "dragonfly" "freebsd" "freebsd" "freebsd" "linux" "linux" "linux" "linux" "linux" "linux" "linux" "linux" "linux" "linux" "netbsd" "netbsd" "netbsd" "openbsd" "openbsd" "openbsd" "plan9" "plan9" "solaris" "windows" "windows")
ARCH=("arm" "386" "amd64" "arm" "arm64" "amd64" "386" "amd64" "arm" "386" "amd64" "arm" "arm64" "ppc64" "ppc64le" "mips" "mipsle" "mips64" "mips64le" "386" "amd64" "arm" "386" "amd64" "arm" "386" "amd64" "amd64" "386" "amd64")

for i in "${!OS[@]}"; do
    os="${OS[$i]}"
    arch="${ARCH[$i]}"
    mkdir -p build/"$os-$arch"
    printf "build.. $os-$arch\n"
    env GOOS="$os" GOARCH="$arch" go build -o build/"$os-$arch"/play github.com/deryrahman/player-cli
done
