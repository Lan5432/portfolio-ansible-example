#!/bin/bash
docker-compose -f ProxyDocker/docker-compose.yml up -d --force-recreate $1
ansible-playbook ansible/proxied_playbook.yml -i ansible/inventory.yml --tags $2