#!/bin/bash
docker run -it --rm --security-opt=seccomp:unconfined -p 8000:3000 -v ~/work/golang/go-learn:/go/src/go-learn -w /go/src/go-learn tour