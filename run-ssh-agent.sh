#!/bin/bash
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/docker_id_rsa