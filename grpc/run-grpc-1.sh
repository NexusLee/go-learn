#!/bin/bash
docker run -it --rm -p 50051:50051 -v ~/work/golang/go-learn:/go/src/go-learn  -w /go/src/go-learn tour
