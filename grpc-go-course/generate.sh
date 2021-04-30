#!/bin/bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.