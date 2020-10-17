#!/bin/bash -x
run_dev () {
sudo podman pull docker.io/containercraft/golang;
sudo /usr/bin/podman run \
    -it --rm --name go-build \
    --volume $(pwd):/root/dev:z \
    --volume $(pwd)/bin:/tmp/bin:z \
    --entrypoint bash \
  docker.io/containercraft/golang; \
}
run_dev

mkdir -p /root/platform/iac/openshift
cp -f site.yml /root/platform/iac/openshift/ 
