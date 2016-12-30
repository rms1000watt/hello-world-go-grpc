package serve

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rms1000watt/hello-world-go-grpc/pb"
)

type helloWorldServer struct{}

// Hello is the RPC function to satisfy gRPC
func (hws *helloWorldServer) Hello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	greetings := "Hello" + in.FirstName + " " + in.LastName

	response := pb.HelloWorldResponse{
		Greetings: greetings,
	}
	return response, nil
}

func main() {
	port := 3456

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hws := helloWorldServer{}
	pb.RegisterHelloWorldServer(s, hws)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
