#!/bin/bash -x
#gitBranch="$(git branch | awk '/\*/{print $2}')"

#sudo chown $USER ~/.ssh -R

#git stage -A; git commit -m "$@"; git push origin ${gitBranch}

#./build.sh

mkdir -p /root/platform/iac/openshift
cp -f ./site.yml /root/platform/iac/openshift/site.yml 2>/dev/null
