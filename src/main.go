package src

import (
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"errors"

	"path/filepath"

	"github.com/rms1000watt/hello-world-go-grpc/pb"
)

// Config is the configuration for the server
type Config struct {
	Address string
	Logging bool
}

// Server is the server obj for hello world server
type Server struct {
	Config
}

// Hello is the RPC function to satisfy the gRPC service
func (s *Server) Hello(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	greetings := "Hello " + req.FirstName + " " + req.LastName

	res := &pb.HelloWorldResponse{
		Greetings: greetings,
	}

	s.log(req, res)
	return res, nil
}

func (s *Server) log(req *pb.HelloWorldRequest, res *pb.HelloWorldResponse) {
	if s.Config.Logging {
		log.Println("REQUEST", req)
		log.Println("RESPONSE", res)
	}
}

// Serve is the main logic for the "serve" command
func Serve(config Config) {
	certFile, keyFile, err := GetServerCertKeyFiles()
	if err != nil {
		log.Fatalln("Error getting cert or key", err)
	}

	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalln("Error listening", err)
	}

	transportCreds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalln("Error reading in certs", err)
	}
	opts := []grpc.ServerOption{grpc.Creds(transportCreds)}

	grpcServer := grpc.NewServer(opts...)
	s := &Server{
		Config: config,
	}
	pb.RegisterHelloWorldServer(grpcServer, s)

	log.Println("Serving on", config.Address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Error serving", err)
	}
}

// GetServerCertKeyFiles returns the cert and key file paths
func GetServerCertKeyFiles() (string, string, error) {
	var err error
	keyFile := ""
	certFile := ""
	foundKey := false
	foundCert := false
	dirs := []string{"./", "./cert"}

	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return "", "", err
		}

		for _, f := range files {
			fullPath := filepath.Join(dir, f.Name())
			switch f.Name() {
			case "cert.pem", "server.cer":
				certFile = fullPath
				foundCert = true
			case "key.pem", "server.key":
				keyFile = fullPath
				foundKey = true
			}
		}

		if foundCert && foundKey {
			return certFile, keyFile, nil
		}
	}

	if !(foundCert && foundKey) {
		err = errors.New("Could not find key and cert")
	}
	return certFile, keyFile, err
}

// GetCACertFile returns the CA cert file path
func GetCACertFile() (string, error) {
	caCertFileName := "ca.cer"
	dirs := []string{"./", "./cert"}

	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return "", err
		}

		for _, f := range files {
			if f.Name() == caCertFileName {
				return filepath.Join(dir, caCertFileName), nil
			}
		}
	}

	return "", errors.New("Could not find " + caCertFileName)
}
