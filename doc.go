//go:generate echo "Generating Protobuf"
//go:generate protoc --go_out=plugins=grpc:. pb/helloWorld.proto

package main
