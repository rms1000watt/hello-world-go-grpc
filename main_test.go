package main

import (
	"context"
	"log"
	"os"
	"testing"

	"strings"

	"github.com/rms1000watt/hello-world-go-grpc/pb"
	"github.com/rms1000watt/hello-world-go-grpc/src"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	address = "localhost:8081"
	logging = false
)

func TestMain(m *testing.M) {
	config := src.Config{
		Address: address,
		Logging: logging,
	}
	go src.Serve(config)
	os.Exit(m.Run())
}

func TestServer(t *testing.T) {
	host := strings.Split(address, ":")[0]

	caCertFile, err := src.GetCACertFile()
	if err != nil {
		log.Fatalln("Error getting CA Cert")
	}

	transportCreds, err := credentials.NewClientTLSFromFile(caCertFile, host)
	if err != nil {
		t.Fatal("Failed getting cert", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(transportCreds)}

	clientConn, err := grpc.Dial(address, opts...)
	if err != nil {
		t.Fatal("Failed connecting", err)
	}
	client := pb.NewHelloWorldClient(clientConn)
	req := &pb.HelloWorldRequest{
		FirstName: "Ryan",
		LastName:  "Smith",
	}
	res, err := client.Hello(context.TODO(), req)
	if err != nil {
		t.Fatal("Failed sending request", err)
	}

	if valid := validateHello(req, res); !valid {
		t.Fatal("Hello response invalid", err)
	}
}

func validateHello(req *pb.HelloWorldRequest, res *pb.HelloWorldResponse) bool {
	ind := 0
	ind = isNegativeOne(ind, strings.Index(res.Greetings, req.FirstName))
	ind = isNegativeOne(ind, strings.Index(res.Greetings, req.LastName))
	return ind != -1
}

func isNegativeOne(old, new int) int {
	if old == -1 {
		return old
	}
	return new
}
