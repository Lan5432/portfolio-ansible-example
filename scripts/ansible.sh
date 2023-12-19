#!/bin/bash
TAGS=$1
echo ansible-playbook ansible/proxied_playbook.yml -i ansible/inventory.yml ${TAGS:+"--tags" "$TAGS"}
ansible-playbook ansible/proxied_playbook.yml -i ansible/inventory.yml ${TAGS:+"--tags" "$TAGS"}