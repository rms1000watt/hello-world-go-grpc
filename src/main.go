package src

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/rms1000watt/hello-world-go-grpc/pb"
)

type helloWorldServer struct{}

// Hello is the RPC function to satisfy gRPC
func (hws *helloWorldServer) Hello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	greetings := "Hello" + in.FirstName + " " + in.LastName

	response := &pb.HelloWorldResponse{
		Greetings: greetings,
	}
	return response, nil
}

func Serve() {
	addr := "127.0.0.1:3456"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hws := &helloWorldServer{}
	pb.RegisterHelloWorldServer(s, hws)

	log.Println("Serving on", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
