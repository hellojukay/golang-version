#!usr/bin/env sh
branch="$(git rev-parse --abbrev-ref --short HEAD)"
commit="$(git rev-parse --short HEAD)"
go build -ldflags "-X main._version=${branch}-${commit}"