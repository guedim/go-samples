#!/bin/bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# gen code from protocol buffer file
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

# init go module
go mod init grpc-greeting 

# look for dependencies
go mod tidy 



