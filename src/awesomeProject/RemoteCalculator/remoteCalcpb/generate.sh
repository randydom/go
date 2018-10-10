#!/usr/bin/env bash

/usr/local/bin/protoc -I ./ remoteCalc.proto --go_out=plugins=grpc:./