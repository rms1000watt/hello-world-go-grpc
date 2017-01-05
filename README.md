<hidden cmd="git push https://rms1000watt@github.com/rms1000watt/hello-world-go-grpc.git master:master"/>
# Hello World Go gRPC

- Cobra (init.sh & cmd)
- gRPC (init.sh & pb & src/main.go)
- GoVendor (init.sh & vendor)
- Docker (doc.go & Dockerfile)
- Flag || Env configurated (cmd/serve.go)
- Tests (main_test.go)
- Client Example (main_test.go)
- Go Generate (doc.go)

## Installation

View `init.sh`

## Usage

```
# Courtesy of https://github.com/deckarep/EasyCert
# Generate CA key, cert & Server key, csr & Sign csr with CA creds
mkdir cert
cd cert
openssl genrsa -out ca.key 2048
openssl req -x509 -new -key ca.key -out ca.cer -days 90 -subj /CN="RMS1000WATT Certificate Authority"
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj /CN="127.0.0.1"
openssl x509 -req -in server.csr -out server.cer -CAkey ca.key -CA ca.cer -days 90 -CAcreateserial -CAserial serial
cd ..

# Run Everything: protoc, go test, go build, docker build
go generate
docker run rms1000watt/hello-world-go-grpc
```

## Development

GoVendor more libs
```
govendor fetch https://github.com/pkg/errors
```

Update protobuf
```
protoc --go_out=plugins=grpc:. pb/helloWorld.proto
```

## TODO

- [] TLS self-signed
- [] TLS letsencrypt (autocert)