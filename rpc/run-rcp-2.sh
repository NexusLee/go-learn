#!/bin/bash
docker run -it --rm --security-opt=seccomp:unconfined -v ~/work/golang/go-learn:/go/src/go-learn -w /go/src/go-learn tour