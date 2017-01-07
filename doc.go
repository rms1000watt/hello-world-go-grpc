//go:generate echo "(You can pass in ENV variables to this command `KEY1=value1 KEY2=value2 go generate`)"
//go:generate echo "Generating Protobuf"
//go:generate protoc --go_out=plugins=grpc:. pb/helloWorld.proto
//go:generate echo "Running Tests"
//go:generate go test
//go:generate echo "Building Linux"
//go:generate sh -c "GOOS=linux go build -o ./bin/hello-world-go-grpc-linux"
//go:generate echo "Dockerizing (Make sure Docker is running)"
//go:generate docker build -t docker.io/rms1000watt/hello-world-go-grpc:latest .
//go:generate echo "(You can run image by executing: `docker run docker.io/rms1000watt/hello-world-go-grpc:latest`)"
//go:generate echo "(You can push image by executing: `docker push docker.io/rms1000watt/hello-world-go-grpc:latest`)"

package main
