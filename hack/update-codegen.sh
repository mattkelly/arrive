#!/usr/bin/env bash

SCRIPT_NAME=$(basename $0)

SRC_DIR=$GOPATH/src/github.com/mattkelly/arrive

go get -u k8s.io/gengo
rm -rf ./vendor/github.com/golang/glog
rm -rf ./vendor/github.com/spf13/pflag

cd $SRC_DIR

vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/mattkelly/arrive/pkg/client github.com/mattkelly/arrive/pkg/apis \
  "arrive:v1alpha1"
