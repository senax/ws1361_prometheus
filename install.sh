#!/bin/sh
sudo cp -f ws1361_prometheus /usr/sbin/ws1361
sudo cp -f ws1361.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now ws1361

