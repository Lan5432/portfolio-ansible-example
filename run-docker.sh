#!/bin/bash
docker build -t target ./Docker
docker rm $(docker stop $(docker ps -a -q --filter name=target --format="{{.ID}}"))
docker run -d -p 2022:22 -p 80:9090 --name target-1 target
docker run -d -p 2023:22 -p 9091:9091 -p 9092:9092 --name target-2 target