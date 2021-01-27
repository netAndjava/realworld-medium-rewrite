#!/bin/bash

echo "$# arguments"
protoPath="source_relative"
out="../../service/api/"

if [ -n "$1" ]; then
    echo "input proto $1"
else
    echo "please input proto file"
fi

if [ -n "$2" ]; then
   echo "input out_path $2" 
   out=$2
fi

if [ -n "$3" ]; then
   echo "input proto_path $3" 
   protoPath=$3
fi

echo "protoc --go-grpc_opt=paths=$protoPath --go-grpc_out=$out $1"
protoc --proto_path=./ --go-grpc_opt=paths=$protoPath --go-grpc_out=$out $1
