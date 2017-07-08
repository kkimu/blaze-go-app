#!/bin/sh
docker stop `docker ps -q -f name=running-go-app`
docker run --rm -d -p 8000:8000 -v $PWD:/go/src/github.com/kkimu/blaze-go-app --link running-mysql:mysql --name running-go-app go-app
