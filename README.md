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
# Gen self signed certs
mkdir cert
curl https://golang.org/src/crypto/tls/generate_cert.go\?m\=text > cert/generate_cert.go
cd cert && go run generate_cert.go --host 127.0.0.1 && cd ..

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