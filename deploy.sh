#!/bin/bash

#if [ "$(id -u)" -ne 0 ]; then
#    echo "Please run this script with sudo or as root."
#    exit 1
#fi

mkdir bin
go build -o bin/label-printer . || exit 1

sudo mkdir /usr/local/label-printer
sudo cp bin/label-printer /usr/local/label-printer/ || exit 1

SERVICE_NAME=label-printer.service

#sudo systemctl stop "${SERVICE_NAME}"

sudo cp "${SERVICE_NAME}" /etc/systemd/system/ || exit 1

sudo systemctl daemon-reload || exit 1

sudo systemctl enable "${SERVICE_NAME}" || exit 1

sudo systemctl start "${SERVICE_NAME}" || exit 1

#sudo systemctl status "${SERVICE_NAME}"
