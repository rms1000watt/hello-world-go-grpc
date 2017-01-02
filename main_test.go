package main

import (
	"context"
	"testing"
	"time"

	"strings"

	"github.com/rms1000watt/hello-world-go-grpc/pb"
	"github.com/rms1000watt/hello-world-go-grpc/src"
	"google.golang.org/grpc"
)

var (
	address = ":8081"
	logging = true
)

func TestMain(m *testing.M) {
	testCompleteCh := make(chan bool)
	serverStoppedCh := make(chan bool)

	// Start server in go routine
	go func(testCompleteCh, serverStoppedCh chan bool) {
		// Start server
		config := src.Config{
			Address: address,
			Logging: logging,
		}
		src.Serve(config)
		<-testCompleteCh
		// Stop server
		serverStoppedCh <- true
	}(testCompleteCh, serverStoppedCh)

	// Give the server some time to start up..
	time.Sleep(5 * time.Second)

	// Run Test
	m.Run()

	// Stop server
	testCompleteCh <- true

	// Wait for server to stop
	<-serverStoppedCh
}

func TestServer(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

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
	ind := -1
	ind = strings.Index(res.Greetings, req.FirstName)
	ind = strings.Index(res.Greetings, req.LastName)
	return ind != -1
}
