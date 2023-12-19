#!/bin/bash
comebackFolder=$(pwd)
cd ./Ansible/roles/app_server/files/golang
go build -o ./golang-app ./main/
cd $comebackFolder