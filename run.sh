#!/bin/sh
docker stop `docker ps -q`
docker run --rm -d -p 8000:8000 --name runnig-go-app go-app
