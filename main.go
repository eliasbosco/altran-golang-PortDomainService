package main

import (
	pb "./portsgrpc"
	"./types"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPortsDbServer
}

func main() {
	config := types.SetupConfig()

	lis, err := net.Listen("tcp", ":" + config.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPortsDbServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
