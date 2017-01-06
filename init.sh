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

# Run cobra
go run vendor/github.com/spf13/cobra/cobra/main.go init
go run vendor/github.com/spf13/cobra/cobra/main.go add serve
# Edit cmd/serve.go
# Edit cmd/root.go

# Create a protobuf directory and add protobuf file
mkdir pb
touch pb/helloWorld.proto
# Edit pb/helloWorld.proto

# Generate go-proto files
protoc --go_out=plugins=grpc:. pb/helloWorld.proto

# Create a bin directory
mkdir bin

# Create a src directory
mkdir src
touch src/main.go
# Edit src/main.go

# Create a dockerfile
cat <<EOF > Dockerfile
FROM scratch
COPY ./bin/hello-world-go-grpc-linux /
COPY ./cert /
EXPOSE 8081
ENTRYPOINT ["/hello-world-go-grpc-linux", "serve"]
EOF

# Courtesy of https://github.com/deckarep/EasyCert
# Generate CA key, cert & Server key, csr & Sign csr with CA creds
# TODO: touch generate-certs.sh
mkdir cert
cd cert
openssl genrsa -out ca.key 2048
openssl req -x509 -new -key ca.key -out ca.cer -days 90 -subj /CN="RMS1000WATT Certificate Authority"
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj /CN="localhost"
# openssl req -new -key server.key -out server.csr -subj '/CN=localhost/subjectAltName=IP.1=127.0.0.1,IP.2=0.0.0.0'
# openssl req -new -key server.key -out server.csr -subj '/CN=localhost/subjectAltName=IP.1=127.0.0.1,IP.2=0.0.0.0' -reqexts SAN 
# openssl req -new -key server.key -out server.csr -subj "/CN=localhost" -reqexts SAN -config <(cat /etc/ssl/openssl.cnf <(printf "[SAN]\nsubjectAltName=IP:127.0.0.1,IP:0.0.0.0")) 
openssl x509 -req -in server.csr -out server.cer -CAkey ca.key -CA ca.cer -days 90 -CAcreateserial -CAserial serial
cd ..


# Create a doc.go file with go:generate 
cat <<EOF > doc.go
//go:generate echo "(You can pass in ENV variables to this command `KEY1=value1 KEY2=value2 go generate`)"
//go:generate echo "Generating Protobuf"
//go:generate protoc --go_out=plugins=grpc:. pb/helloWorld.proto
//go:generate echo "Generating Certs"
//go:generate echo "TODO: bash generate-certs.sh"
//go:generate echo "Running Tests"
//go:generate go test
//go:generate echo "Building Linux"
//go:generate sh -c "GOOS=linux go build -o ./bin/hello-world-go-grpc-linux"
//go:generate echo "Dockerizing"
//go:generate docker build -t docker.io/rms1000watt/hello-world-go-grpc:latest .
//go:generate echo "(You can run image by executing: `docker run docker.io/rms1000watt/hello-world-go-grpc:latest`)"
//go:generate echo "(You can push image by executing: `docker push docker.io/rms1000watt/hello-world-go-grpc:latest`)"

package main
EOF

# Create a .gitignore
cat <<EOF > .gitignore
.DS_Store
hello-world-go-grpc*
EOF

# Create tests
touch main_test.go
# Edit main_test.go

