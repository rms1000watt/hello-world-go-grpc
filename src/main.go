package src

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/rms1000watt/hello-world-go-grpc/pb"
)

// Config is the configuration for the server
type Config struct {
	Address string
	Logging bool
}

type server struct {
	config Config
}

// Hello is the RPC function to satisfy the gRPC service
func (s *server) Hello(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	greetings := "Hello " + req.FirstName + " " + req.LastName

	res := &pb.HelloWorldResponse{
		Greetings: greetings,
	}

	s.log(req, res)
	return res, nil
}

func (s *server) log(req *pb.HelloWorldRequest, res *pb.HelloWorldResponse) {
	if s.config.Logging {
		log.Println("REQUEST", req)
		log.Println("RESPONSE", res)
	}
}

// Serve is the main logic for the "serve" command
func Serve(config Config) {
	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := &server{
		config: config,
	}
	pb.RegisterHelloWorldServer(grpcServer, s)

	log.Println("Serving on", config.Address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
