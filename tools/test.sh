#!/bin/bash -x
#gitBranch="$(git branch | awk '/\*/{print $2}')"

#sudo chown $USER ~/.ssh -R

#git stage -A; git commit -m "$@"; git push origin ${gitBranch}

#./build.sh

mkdir -p /root/deploy/ansible/deploy
cp -f ./site.yml /root/deploy/ansible/deploy/site.yml 2>/dev/null
