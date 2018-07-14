#!/bin/bash

PLATFORMS="darwin/amd64 linux/amd64"

SCRIPT_NAME=$(basename "${0}")
FAILURES=""

if [ -z "${GOPATH}" ]; then
    GOPATH="${HOME}/go"
fi
BIN_PATH="${GOPATH}/src/github.com/willjcj/param/bin"

for PLATFORM in ${PLATFORMS}; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_FILENAME="${BIN_PATH}/param-${GOOS}-${GOARCH}"

  CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILENAME} ${@}"
  echo "${CMD}"
  eval "${CMD}" || FAILURES="${FAILURES} ${PLATFORM}"
done

# eval errors
if [[ "${FAILURES}" != "" ]]; then
  echo ""
  echo "${SCRIPT_NAME} failed on: ${FAILURES}"
  exit 1
fi
