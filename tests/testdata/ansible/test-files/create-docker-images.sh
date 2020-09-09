#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="wanliuno/cppjit-testrunner" docker-cppjit
docker build --tag="wanliuno/python-testrunner" docker-python
docker build --tag="wanliuno/go-testrunner" docker-go
