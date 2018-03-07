#!/bin/bash

IMAGE=tour
docker build -t ${IMAGE} .
#docker build --build-arg=[HTTP_PROXY=http://192.168.200.163:1080,HTTPS_PROXY=http://192.168.200.163:1080] -t ${IMAGE} .