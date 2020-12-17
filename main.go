package main

import (
	pb "./portsgrpc"
	"./types"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"log"
	"net"
	"./db"
	dbTypes "./db/types"
)

var (
	ports = []string{":10001", ":10002", ":10003"}
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPortsDbServer
}

func (s *server) Upset(ctx context.Context, in *pb.Ports) (*pb.Response, error) {
	log.Printf("Received port %s", in.PortId)

	config := types.SetupConfig()
	log.Printf("main.Upset-config: %v\n", config)

	db.SetConfig(&config)
	_db, err := db.Connect()
	if err != nil {
		log.Printf("main.Upset: %v\n", err)
		return &pb.Response{Code: "sql", Message: err.Error()}, err
	}

	msg := "Port saved in the data base successfully."
	if _, err := _db.Upset(dbTypes.Ports{
		PortId:      in.PortId,
		Name:        in.Name,
		City:        in.City,
		Country:     in.Country,
		Alias:       in.Alias,
		Regions:     in.Regions,
		Coordinates: in.Coordinates,
		Province:    in.Province,
		Timezone:    in.Timezone,
		Unlocs:      in.Unlocs,
		Code:        in.Code,
	}); err != nil {
		log.Printf("main.Upset: %v\n", err)
		msg = err.Error()
	}

	return &pb.Response{Code: "portsgrpc.Upset-1", Message: msg}, nil
}

func main() {
	config := types.SetupConfig()

	lis, err := net.Listen("tcp", ":" + config.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	service.RegisterChannelzServiceToServer(s)
	go s.Serve(lis)
	defer s.Stop()

	/***** Start three GreeterServers(with one of them to be the slowServer). *****/
	for i := 0; i < 3; i++ {
		lis, err := net.Listen("tcp", ports[i])
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()
		s := grpc.NewServer()
		pb.RegisterPortsDbServer(s, &server{})
		go s.Serve(lis)
	}

	/***** Wait for user exiting the program *****/
	select {}
}
