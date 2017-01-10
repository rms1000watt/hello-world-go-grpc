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

# Create an openssl config
# EDIT the 
cat <<EOF > openssl.cnf
[ req ]
distinguished_name = req_distinguished_name
req_extensions     = v3_req

[ req_distinguished_name ]
commonName         = Common Name (e.g. server FQDN or YOUR name)
commonName_default = localhost
commonName_max     = 64

[ v3_req ]
basicConstraints = CA:FALSE
keyUsage         = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName   = @alt_names

[ alt_names ]
DNS.1 = www.example.com
DNS.2 = localhost
IP.1  = 127.0.0.1
IP.2  = 0.0.0.0
EOF

# Courtesy of https://github.com/deckarep/EasyCert
# Generate CA key, cert & Server key, csr & Sign csr with CA creds
mkdir cert
cd cert
openssl genrsa -out ca.key 2048
openssl req -x509 -new -key ca.key -out ca.cer -days 90 -subj /CN="rms1000watt Certificate Authority"
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config ../openssl.cnf
openssl x509 -req -in server.csr -out server.cer -CAkey ca.key -CA ca.cer -days 90 -CAcreateserial -CAserial serial
cd ..

# Create a .gitignore
cat <<EOF > .gitignore
.DS_Store
hello-world-go-grpc*
EOF

# Create tests
touch main_test.go
# Edit main_test.go
