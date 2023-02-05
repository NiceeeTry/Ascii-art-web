#!/bin/bash

docker image build -f Dockerfile -t ascii-art-web-docker .
docker container run -p 8080:8080 --name ascii-art-web ascii-art-web-docker
docker ps -a
# docker rm $(docker ps -qa)