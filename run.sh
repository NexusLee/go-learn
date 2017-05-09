#!/bin/bash
docker run -it --rm -p 8000:3000 -v ~/work/golang/learn:/go/src/learn -w /go/src/learn tour