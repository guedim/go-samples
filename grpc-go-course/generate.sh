#!/bin/bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# init go module
go mod init grpc-greeting 

# look for dependencies
go mod tidy 

#
#   GREETING
#
# gen code from protocol buffer file
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
# run server
go run greet/greet_server/server.go
# fun client
go run greet/greet_client/client.go



#
#   SUM
#
# gen code from protocol buffer file
protoc calculator/calcpb/sum.proto --go_out=plugins=grpc:.
# run server
go run calculator/calc_server/server.go
# fun client
go run calculator/calc_client/client.go


