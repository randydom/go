#!/usr/bin/env bash

/usr/local/bin/protoc -I ./ fileShare.proto --go_out=plugins=grpc:./