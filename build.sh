#!/bin/bash
# cobra init --pkg-name github.com/containercraft/konductor-go
# cobra add mirror
# cobra add bundle
# go build
# gitup devel

goCmd=$(which go)

rm /bin/konductor 2>/dev/null
rm -rf /root/konductor 2>/dev/null
mkdir -p /tmp/bin

plugins="
    github.com/containercraft/konductor-go/cmd \
"
#   "github.com/containercraft/konductor-go/plugins/err" \
#   "github.com/containercraft/konductor-go/plugins/log" \
#   "github.com/containercraft/konductor-go/plugins/auth" \
#   "github.com/mitchellh/mapstructure" \
#   "golang.org/x/sys/unix" \
#   "github.com/spf13/cobra" \
#   "github.com/spf13/viper" \
#   "github.com/go-git/go-git" \
#   "github.com/go-git/go-git/plumbing" \

for i in ${plugins}; do
  ${goCmd} get -u ${i};
done

${goCmd} build

cp -f ./dev /tmp/bin/konductor 2>/dev/null
ls -lah /tmp/bin

