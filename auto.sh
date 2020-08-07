#!/bin/bash -x
run$(\
sudo /usr/bin/podman run \
    -it --rm --name go-build \
    --volume $(pwd)/bin:/tmp/bin:z \
    --entrypoint /root/dev/build.sh \
    --volume $(pwd):/root/dev:z \
  docker.io/ocpredshift/red-gotools; \
  exit 0)
    
ls -lah $(pwd)
ls -lah $(pwd)/bin
exit 0
