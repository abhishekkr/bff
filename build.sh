#!/usr/bin/env bash

set -ex

PROJECT_DIR=$(dirname $0)
BUILD_DIR="${PROJECT_DIR}/out"

buildBinaries(){
  [[ ! -d "${BUILD_DIR}" ]] && mkdir "${BUILD_DIR}"

  set -ex

  go build -o "${BUILD_DIR}/bff" main.go

  GOOS=linux go build -o "${BUILD_DIR}/bff-linux-amd64" main.go

  GOOS=windows go build -o "${BUILD_DIR}/bff-windows-amd64" main.go

  GOOS=darwin go build -o "${BUILD_DIR}/bff-darwin-amd64" main.go

  set +ex
}

##### main
buildBinaries
