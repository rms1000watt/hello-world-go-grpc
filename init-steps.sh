# Assuming golang is installed, else install with Homebrew and define GOPATH

# Download protobuf
# https://github.com/google/protobuf/releases/tag/v3.0.0
# place bin at: /usr/local/bin/protoc

# Download protobuf compiler/generator for go
go get -u github.com/golang/protobuf/protoc-gen-go

# Create project path
mkdir hello-world-go-grpc
cd hello-world-go-grpc

# Install and run govendor
go get -u github.com/kardianos/govendor
govendor init
govendor fetch google.golang.org/grpc
govendor fetch google.golang.org/grpc/reflection

# Install globally and run cobra
go get -v github.com/spf13/cobra/cobra
cobra init
cobra add serve

# Create a protobuf directory and add protobuf file
mkdir pb
touch pb/helloWorld.proto
# Edit helloWorld.proto

# Generate go-proto files
protoc --go_out=plugins=grpc:. pb/helloWorld.proto


# Create a src directory
mkdir src
mkdir src/serve
touch src/serve/main.go
# Edit main.go


