#!/usr/bin/env bash

mkdir -p /opt/varnost/sidecar/
make linux
mv bin/varnost-sidecar* /opt/varnost/sidecar/
cp vanrost-sidecar.service /etc/systemd/system/varnost-sidecar.service
systemctl daemon-reload
systemctl start varnost-sidecar
systemctl enable varnost-sidecar
