#!/bin/bash

parent_dir=$(dirname "$PWD")

cd parent_dir

mkdir bin

go run parent_dir || exit 1

sudo mkdir /usr/local/label-printer

sudo cp bin/label-printer /usr/local/label-printer/ || exit 1

SERVICE_NAME=label-printer.service

sudo cp "${SERVICE_NAME}" /etc/systemd/system/ || exit 1

sudo systemctl daemon-reload || exit 1

sudo systemctl enable "${SERVICE_NAME}" || exit 1

sudo systemctl start "${SERVICE_NAME}" || exit 1
