#!/bin/sh
VERSION=${1:-dev}
docker build -f Dockerfile-$VERSION -t wxwmatt/go-service:1.0 .
