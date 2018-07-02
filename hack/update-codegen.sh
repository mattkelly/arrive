#!/usr/bin/env bash

set -o errexit
set -o nounset
# TODO uncomment when gengo workaround is removed
#set -o pipefail

# TODO remove these workarounds
go get -u k8s.io/gengo || true
rm -rf ./vendor/github.com/golang/glog
rm -rf ./vendor/github.com/spf13/pflag

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

${CODEGEN_PKG}/generate-groups.sh all \
  github.com/mattkelly/arrive/pkg/client github.com/mattkelly/arrive/pkg/apis \
  "arrive:v1alpha1"
