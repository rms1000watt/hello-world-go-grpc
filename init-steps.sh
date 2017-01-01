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

# Generate go-proto files
protoc --go_out=plugins=grpc:. pb/helloWorld.proto

# Create a src directory
mkdir src
touch src/main.go
# Edit src/main.go

# Create a dockerfile
cat <<EOF > Dockerfile
FROM scratch
COPY ./hello-world-go-grpc-linux /
EXPOSE 8081
ENTRYPOINT ["/hello-world-go-grpc-linux", "serve"]
EOF

# Create a doc.go file with go:generate 
cat <<EOF > doc.go
//go:generate echo "(You can pass in ENV variables to this command `KEY1=value1 KEY2=value2 go generate`)"
//go:generate echo "Generating Protobuf"
//go:generate protoc --go_out=plugins=grpc:. pb/helloWorld.proto
//go:generate echo "Building Linux"
//go:generate sh -c "GOOS=linux go build -o hello-world-go-grpc-linux"
//go:generate echo "Dockerizing"
//go:generate docker build -t docker.io/rms1000watt/hello-world-go-grpc:latest .
//go:generate echo "(You can push repo by running: `docker push docker.io/rms1000watt/hello-world-go-grpc:latest`)"

package main
EOF



# Create a .gitignore
cat <<EOF > .gitignore
.DS_Store
hello-world-go-grpc*
EOF
