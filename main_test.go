package main

import (
	"context"
	"log"
	"net"
	"testing"

	"strings"

	"github.com/rms1000watt/hello-world-go-grpc/pb"
	"github.com/rms1000watt/hello-world-go-grpc/src"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	address  = "127.0.0.1:8081"
	logging  = false
	certFile = ""
	keyFile  = ""
)

func TestMain(m *testing.M) {
	var err error
	doneCh := make(chan bool)

	// Define config
	config := src.Config{
		Address: address,
		Logging: logging,
	}

	// Get certs
	certFile, keyFile, err = src.GetCertKeyFiles()
	if err != nil {
		log.Fatalln("Error getting cert or key", err)
	}

	// Get listener
	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalln("Error listening", err)
	}

	// Start server in goroutine
	go func(doneCh chan bool) {
		// Setup TLS
		transportCreds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalln("Error reading in certs", err)
		}
		opts := []grpc.ServerOption{grpc.Creds(transportCreds)}

		// Start server
		grpcServer := grpc.NewServer(opts...)
		s := &src.Server{
			Config: config,
		}
		pb.RegisterHelloWorldServer(grpcServer, s)

		// Ignore error since it WILL error because lis.Close() called below
		grpcServer.Serve(lis)

		// For proper cleanup
		doneCh <- true
	}(doneCh)

	// Give the server some time to start up..
	// time.Sleep(3 * time.Second)

	// Run Test
	m.Run()

	// Close listener
	err = lis.Close()
	if err != nil {
		log.Println("Error closing listener", err)
	}

	// For proper cleanup
	<-doneCh
}

func TestServer(t *testing.T) {
	host := strings.Split(address, ":")[0]
	transportCreds, err := credentials.NewClientTLSFromFile(certFile, host)
	if err != nil {
		t.Fatal("Failed getting cert", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(transportCreds)}

	// TODO: fix bad sever CA-signed-cert (because of self signed cert issue)
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
