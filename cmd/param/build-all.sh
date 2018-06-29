#!/bin/bash
#
# Trimmed down version of: https://gist.github.com/eduncan911/68775dba9d3c028181e4
#
# GoLang cross-compile snippet for Go 1.6+ based loosely on Dave Chaney's cross-compile script:
# http://dave.cheney.net/2012/09/08/an-introduction-to-cross-compilation-with-go
#
# To use:
#
#   $ cd ~/path-to/my-awesome-project
#   $ go-build-all
#
PLATFORMS="darwin/amd64 linux/amd64"

type setopt >/dev/null 2>&1

SCRIPT_NAME=`basename "$0"`
FAILURES=""
SOURCE_FILE=`echo $@ | sed 's/\.go//'`
CURRENT_DIRECTORY=${PWD##*/}
OUTPUT=${SOURCE_FILE:-$CURRENT_DIRECTORY} # if no src file given, use current dir name

for PLATFORM in $PLATFORMS; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_FILENAME="${OUTPUT}-${GOOS}-${GOARCH}"
  if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
  CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILENAME} $@"
  echo "${CMD}"
  eval $CMD || FAILURES="${FAILURES} ${PLATFORM}"
done

# eval errors
if [[ "${FAILURES}" != "" ]]; then
  echo ""
  echo "${SCRIPT_NAME} failed on: ${FAILURES}"
  exit 1
fi
