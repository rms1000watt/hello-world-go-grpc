# Assuming golang is installed, else install with Homebrew and define GOPATH

# Download protobuf
# https://github.com/google/protobuf/releases/tag/v3.0.0
# place bin at: /usr/local/bin/protoc

# Download protobuf compiler/generator for go
# TODO: vendor this as well
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# Create project path
mkdir hello-world-go-grpc
cd hello-world-go-grpc

# Install and run govendor
go get -u github.com/kardianos/govendor
govendor init
govendor fetch google.golang.org/grpc
govendor fetch github.com/spf13/cobra/cobra

# run cobra
go run vendor/github.com/spf13/cobra/cobra/main.go init
go run vendor/github.com/spf13/cobra/cobra/main.go add serve

# Create a protobuf directory and add protobuf file
mkdir pb
touch pb/helloWorld.proto
# Edit pb/helloWorld.proto

# Create a doc.go file with go:generate 
cat <<EOF > doc.go
//go:generate echo "Generating Protobuf"
//go:generate protoc --go_out=plugins=grpc:. pb/helloWorld.proto

package main
EOF

# Generate go-proto files
go generate

# Create a src directory
mkdir src
touch src/main.go
# Edit src/main.go



