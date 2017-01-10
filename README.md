<hidden cmd="git push https://rms1000watt@github.com/rms1000watt/hello-world-go-grpc.git master:master"/>

# Hello World Go gRPC

- Cobra                     (init.sh & cmd)
- gRPC                      (init.sh & pb & src/main.go)
- GoVendor                  (init.sh & vendor)
- Docker                    (doc.go & Dockerfile)
- Flag || Env configurated  (cmd/serve.go)
- Tests                     (main_test.go)
- Client Example            (main_test.go)
- Go Generate               (doc.go)
- TLS Support               (init.sh & openssl.cnf)

## Build Repo from Scratch

View `init.sh`

## Usage

### Generate Certs

```
# Courtesy of https://github.com/deckarep/EasyCert
# Generate CA key, cert & Server key, csr & Sign csr with CA creds
# UPDATE openssl.cnf with proper DNS & IP SANs
mkdir cert
cd cert
openssl genrsa -out ca.key 2048
openssl req -x509 -new -key ca.key -out ca.cer -days 90 -subj /CN="rms1000watt Certificate Authority"
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config ../openssl.cnf
openssl x509 -req -in server.csr -out server.cer -CAkey ca.key -CA ca.cer -days 90 -CAcreateserial -CAserial serial
cd ..
```

Useful command to verify 

```
openssl req -text -noout -in cert/server.csr
```

### Update Protobuf

```
protoc --go_out=plugins=grpc:. pb/helloWorld.proto
```

### Run Tests

```
go test
```

### Build Binary

```
GOOS=linux go build -o ./bin/hello-world-go-grpc-linux
```

### Build Docker Image

```
docker build -t docker.io/rms1000watt/hello-world-go-grpc:latest .
```

### Run Docker Image

```
docker run docker.io/rms1000watt/hello-world-go-grpc:latest
```

Environmental variables can be passed in to container

```
docker run -e ADDRESS=127.0.0.1:8081 -e LOGGING=true docker.io/rms1000watt/hello-world-go-grpc:latest
```

### Push Docker Image

```
docker push docker.io/rms1000watt/hello-world-go-grpc:latest
```

## Continued Development

GoVendor more libs...

```
govendor fetch https://github.com/pkg/errors
```

...and continue coding. Update the protobuf as necessary.
