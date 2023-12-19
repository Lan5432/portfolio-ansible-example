#!/bin/bash
docker-compose -f ProxyDocker/docker-compose.yml up -d --force-recreate $1
source scripts/ansible.sh $2