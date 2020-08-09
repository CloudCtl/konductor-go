#!/bin/bash

rm -f dev 2>/dev/null
sudo chown $USER ~/.ssh -R
git stage -A; git commit -m 'testing'; git push origin master
./build.sh
cp -f ./bin/konductor /usr/bin/konductor 2>/dev/null
cp -f ./dev /usr/bin/konductor 2>/dev/null
