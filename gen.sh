#!/bin/bash
gowrap gen -p io -i Reader -t ./templates/ratelimit -o generated.go &&
protoc -I grpc/ grpc/*.proto --go_out=plugins=grpc:grpc

