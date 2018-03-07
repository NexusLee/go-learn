#!/bin/bash
docker run -it --rm -p 8001:3001 -v ~/work/golang/go-learn:/go/src/go-learn  -w /go/src/go-learn tour
