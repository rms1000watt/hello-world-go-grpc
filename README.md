<hidden cmd="git push https://rms1000watt@github.com/rms1000watt/hello-world-go-grpc.git master:master"/>
# Hello World Go gRPC

- Cobra (init-steps.sh & cmd)
- gRPC (init-steps.sh & pb & src/main.go)
- GoVendor (init-steps.sh & vendor)
- Docker (doc.go & Dockerfile)
- Flag || Env configurated (cmd/serve.go)
- Tests (main_test.go)
- Client Example (main_test.go)
- Go Generate (doc.go)

## Installation

View `init-steps.sh`

## Usage

Update protobuf
```
protoc --go_out=plugins=grpc:. pb/helloWorld.proto
```

Run everything
```
# Runs: protoc, go test, go build, docker build
go generate
docker run rms1000watt/hello-world-go-grpc
```

## TODO

- [] TLS