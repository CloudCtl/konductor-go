#!/bin/bash
# cobra init --pkg-name github.com/CodeSparta/konductor-go
# cobra add mirror
# cobra add bundle
# go build
# gitup devel

goCmd=$(which go)

rm /bin/konductor 2>/dev/null
rm -rf /root/konductor 2>/dev/null
mkdir -p /tmp/bin

sudo chown $USER ~/.ssh -R
git stage -A; git commit -m 'testing'; git push origin master

plugins="
    "golang.org/x/sys/unix" \
    "github.com/spf13/viper" \
    "github.com/spf13/cobra" \
    "github.com/go-git/go-git" \
    "github.com/go-git/go-git/plumbing" \
    "github.com/CodeSparta/konductor-go/cmd" \
    "github.com/CodeSparta/konductor-go/plugins/err" \
    "github.com/CodeSparta/konductor-go/plugins/log" \
    "github.com/CodeSparta/konductor-go/plugins/auth" \
"
for i in ${plugins}; do
  ${goCmd} get -u ${i};
done

${goCmd} build

mv ./dev ./bin/konductor 2>/dev/null
cp -f ./bin/konductor /usr/bin/konductor 2>/dev/null
cp -f ./bin/konductor /tmp/bin/konductor 2>/dev/null

