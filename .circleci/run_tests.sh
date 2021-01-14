#!/bin/bash

set -e

export GOROOT=/usr/local/go
export GOPATH=/tmp/gopath
export PATH=$PATH:$GOROOT/bin

mkdir -p $GOPATH/src/github.com/benderdave
cp -r /home/circleci/project/ $GOPATH/src/github.com/benderdave/fsquota
cd $GOPATH/src/github.com/benderdave/fsquota

export TEST_MOUNTPOINT_QUOTAS_ENABLED=/mnt/quota_test
export TEST_MOUNTPOINT_QUOTAS_DISABLED=/mnt/noquota_test
go test -v --covermode=atomic --coverprofile=/home/circleci/project/coverage.txt .
chmod 0644 /home/circleci/project/coverage.txt
