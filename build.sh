#!/bin/bash

docker build -t fakeconsul .
docker run -p 8500:8500 fakeconsul