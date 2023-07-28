#!/bin/bash

set -o errexit
set -o errtrace
set -o nounset
set -o pipefail


if [ -z "$(command -v podman)" ]; then
  echo "ERR: podman is not installed"
  echo "RUN: sudo apt install podman"
  exit 1
fi

podman run --name arangodb --hostname arangodb -e ARANGO_NO_AUTH=1 -p 8529:8529 -d arangodb

