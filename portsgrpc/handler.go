package portsgrpc

import (
	"log"

	"golang.org/x/net/context"
	"../types"
	"../db"
	dbTypes "../db/types"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) Upset(ctx context.Context, in *Ports) (*Response, error) {
	log.Printf("Received port %s", in.PortId)
	config := types.SetupConfig()
	dbConn, err := db.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	msg := "Port saved in the data base successfully."
	if _, err := dbConn.Upset(integrateModels(in)); err != nil {
		log.Fatal(err)
		msg = err.Error()
	}

	return &Response{Code: "portsgrpc.Upset-1", Message: msg}, nil
}

func integrateModels(in *Ports) dbTypes.Ports {
	return dbTypes.Ports{
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
	}
}