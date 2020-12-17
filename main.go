package main

import (
	pb "./portsgrpc"
	"./types"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

var (
	ports = []string{":10001", ":10002", ":10003"}
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPortsDbServer
}

func (s *server) Upsert(ctx context.Context, in *pb.Ports) (*pb.Response, error) {
	log.Printf("Received ports %#v", in.PortsBody)

	config := types.SetupConfig()
	log.Printf("main.Upsert.config: %v\n", config)

	/*
		All the database functions will be put here because of context issues.
		When I've tried to set the context.WithValue variable set the
		database connection object, it has a strange behave, return
		database closed message
	*/
	dns := config.MysqlUsername + ":" + config.MysqlPassword + "@tcp(" + config.MysqlAddr + ")/mysql"
	log.Printf("portsgrpc.Upsert.db.Connect: %s\n", dns)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Printf("portsgrpc.Upsert.db.Connect: %v\n", err)
		return &pb.Response{Code: "portsgrpc.Upsert.db.Connect", Message: err.Error()}, err
	}
	defer db.Close()

	sql := `CREATE TABLE IF NOT EXISTS ports (
		port_id char(5) PRIMARY KEY,
		name varchar(100) NOT NULL,
		city varchar(100) NOT NULL,
		country varchar(100) NOT NULL,
		alias varchar(500),
		regions varchar(500),
		coordinates varchar(100),
		province varchar(100),
		timezone varchar(50),
		unlocs varchar(500),
		code varchar (100)
	);` // SQL Statement for Create Table

	statement, err := db.Prepare(sql) // Prepare SQL Statement
	if err != nil {
		log.Printf("portsgrpc.Upsert.db.createTable: %v\n", err.Error())
		return &pb.Response{Code: "portsgrpc.Upsert.db.createTable", Message: err.Error()}, err
	}
	statement.Exec() // Execute SQL Statements
	defer statement.Close()
	log.Printf("Create table ...\n%v\n", sql)

	for _, portsBody := range in.PortsBody {
		// Verifying whether record exists in database before
		sql = `SELECT port_id FROM ports WHERE port_id = ?`
		row, err := db.Query(sql, portsBody.PortId)
		if err != nil {
			log.Printf("portsgrpc.Upsert.db.Exists: %v\n", err)
			return &pb.Response{Code: "portsgrpc.Upsert.db.Exists", Message: err.Error()}, err
		}
		defer row.Close()
		portId := ""
		for row.Next() { // Iterate and fetch the records from result cursor
			row.Scan(&portId)
		}
		exists := (portId != "")

		tx, err := db.Begin()
		if err != nil {
			log.Printf("portsgrpc.Upsert.db.Upsert: %v\n", err)
			return &pb.Response{Code: "portsgrpc.Upsert.db.Upsert-begin", Message: err.Error()}, err
		}
		if !exists {
			sql := `insert into ports 
		(port_id,name,city,country,alias,regions,coordinates,province,timezone,unlocs,code) 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
			stmt, err := tx.Prepare(sql)
			if err != nil {
				log.Printf("portsgrpc.Upsert.db.Upsert-insert: %v\n", err)
			}
			defer stmt.Close()
			log.Printf("portsgrpc.Upsert.db.Upsert-insert: %s\n", sql)

			_, err = stmt.Exec(portsBody.PortId,
				portsBody.Name,
				portsBody.City,
				portsBody.Country,
				strings.Join(portsBody.Alias, ","),
				strings.Join(portsBody.Regions, ","),
				fmt.Sprintf("%v", portsBody.Coordinates),
				portsBody.Province,
				portsBody.Timezone,
				strings.Join(portsBody.Unlocs, ","),
				portsBody.Code)
			if err != nil {
				log.Printf("portsgrpc.Upsert.db.Upsert-insert: %v\n", err)
			}
		} else {
			sql := `update ports set 
		name=?,city=?,country=?,alias=?,regions=?,coordinates=?,province=?,timezone=?,unlocs=?,code=? 
		where port_id=?`
			stmt, err := tx.Prepare(sql)
			if err != nil {
				log.Printf("portsgrpc.Upsert.db.Upsert-update: %v\n", err)
			}
			defer stmt.Close()
			log.Printf("portsgrpc.Upsert.db.Upsert-update: %s\n", sql)

			_, err = stmt.Exec(portsBody.Name,
				portsBody.City,
				portsBody.Country,
				strings.Join(portsBody.Alias, ","),
				strings.Join(portsBody.Regions, ","),
				fmt.Sprintf("%v", portsBody.Coordinates),
				portsBody.Province,
				portsBody.Timezone,
				strings.Join(portsBody.Unlocs, ","),
				portsBody.Code,
				portsBody.PortId)
			if err != nil {
				log.Printf("portsgrpc.Upsert.db.Upsert-update: %v\n", err)
			}
		}

		if err = tx.Commit(); err != nil {
			log.Printf("portsgrpc.Upsert.db.Upsert-commit: %v\n", err)
		}
	}

	return &pb.Response{
		Code: "portsgrpc.Upsert-finished",
		Message: "Ports saved in the database.",
	}, nil
}

func main() {
	config := types.SetupConfig()
	grpcAddress := strings.Split(config.GrpcAddress, ",")
	log.Printf("main.grpcAddress: %#v\n", grpcAddress)

	/***** Start three GreeterServers(with one of them to be the slowServer). *****/
	for i := 0; i < 3; i++ {
		lis, err := net.Listen("tcp", grpcAddress[i])
		if err != nil {
			log.Fatalf("main.Listen: %v", err)
		}
		defer lis.Close()
		s := grpc.NewServer()
		pb.RegisterPortsDbServer(s, &server{})
		go s.Serve(lis)
	}

	/***** Wait for user exiting the program *****/
	select {}
}
