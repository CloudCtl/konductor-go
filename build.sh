#!/bin/bash -x
# cobra init --pkg-name github.com/containercraft/konductor-go
# cobra add mirror
# cobra add bundle
# go build
# gitup devel

goCmd=$(which go)

rm /bin/konductor 2>/dev/null
rm -rf /root/konductor 2>/dev/null
mkdir -p /tmp/bin

${goCmd} mod download
${goCmd} build

cp -f ./dev /tmp/bin/konductor 2>/dev/null
ls -lah /tmp/bin
./tools/test.sh
./bin/konductor deploy
