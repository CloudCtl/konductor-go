#!/bin/bash
gitBranch="$(git branch | awk '/\*/{print $2}')"

sudo chown $USER ~/.ssh -R

git stage -A; git commit -m "$@"; git push origin ${gitBranch}

./build.sh

cp -f ./bin/konductor /usr/bin/konductor 2>/dev/null
mv -f ./dev /usr/bin/konductor 2>/dev/null
