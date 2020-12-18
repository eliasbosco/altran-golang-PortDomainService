package main

import (
	pb "./portsgrpc"
	"./types"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var (
	ports = []string{":10001", ":10002", ":10003"}
)

// server is used to implement portsgrpc
type server struct {
	pb.UnimplementedPortsDbServer
}

func (s *server) GetPortsDb(ctx context.Context, in *pb.Request) (*pb.Ports, error) {
	log.Printf("Fetch database results with request: %#v ...", in)

	config := types.SetupConfig()

	if _, err := os.Stat(config.SqlitePath); err != nil {
		log.Printf("portsgrpc.Upsert.config: %v - creating new database file\n", err)
		file, err := os.Create(config.SqlitePath) // Create SQLite file
		if err != nil {
			log.Printf("connectSqlite.File: %v\n", err)
			return nil, err
		}
		file.Close()
		log.Printf("connectSqlite.File: Database file '%s' created", config.SqlitePath)
	}
	db, err := sql.Open("sqlite3", config.SqlitePath)
	if err != nil {
		log.Printf("connectSqlite.Open: %v\n", err)
		return nil, err
	}
	defer db.Close()

	//Fetch database results
	var rows *sql.Rows
	sql := "SELECT port_id,name,city,country,alias,regions,coordinates,province,timezone,unlocs,code FROM ports"

	if in.PortId != "" {
		sql += " WHERE port_id = ?"
		rows, err = db.Query(sql, in.PortId)
		if err != nil {
			log.Printf("portsgrpc.GetPortsBody.Query: %v\n", err)
			return &pb.Ports{}, err
		}
	} else {
		sql += " OFFSET " + string(in.Skip) + " LIMIT " + string(in.Limit)
		rows, err = db.Query(sql)
		if err != nil {
			log.Printf("portsgrpc.GetPortsBody.Query: %v\n", err)
			return &pb.Ports{}, err
		}
	}
	defer rows.Close()

	var portsBodyArr []*pb.PortsBody
	for rows.Next() { // Iterate and fetch the records from result cursor
		_portId := ""
		name := ""
		city := ""
		country := ""
		alias := ""
		regions := ""
		coordinates := ""
		province := ""
		timezone := ""
		unlocs := ""
		code := ""
		rows.Scan(&_portId,&name,&city,&country,&alias,&regions,&coordinates,&province,&timezone,&unlocs,&code)

		var _coord []float32
		if coordinates != "" {
			coordArr := strings.Split(coordinates[1:len(coordinates)-1], " ")
			for _, c := range coordArr {
				value, err := strconv.ParseFloat(c, 32)
				if err != nil {
					// do something sensible
				}
				_coord = append(_coord, float32(value))
			}
		}
		portsBodyArr = append(portsBodyArr, &pb.PortsBody{
			PortId: _portId,
			Name: name,
			City: city,
			Country: country,
			Alias: strings.Split(alias, ","),
			Regions: strings.Split(regions, ","),
			Coordinates: _coord,
			Province: province,
			Timezone: timezone,
			Unlocs: strings.Split(unlocs, ","),
			Code: code,
		})
	}

	log.Printf("_ports: %#v\n", &portsBodyArr)
	return &pb.Ports{PortsBody: portsBodyArr}, nil
}

func (s *server) Upsert(ctx context.Context, in *pb.Ports) (*pb.Response, error) {
	log.Printf("Received ports %d records\n", len(in.PortsBody))

	config := types.SetupConfig()
	log.Printf("portsgrpc.Upsert.config: %v\n", config)

	/*
		All the database functions will be implemented here because of context issues.
		When I've tried to set context.WithValue with a variable to the
		database connection object, a strange behave happen, returning
		database closed message.
	*/
	if _, err := os.Stat(config.SqlitePath); err != nil {
		log.Printf("portsgrpc.Upsert.config: %v - creating new database file\n", err)
		file, err := os.Create(config.SqlitePath) // Create SQLite file
		if err != nil {
			log.Printf("connectSqlite.File: %v\n", err)
			return nil, err
		}
		file.Close()
		log.Printf("connectSqlite.File: Database file '%s' created", config.SqlitePath)
	}
	db, err := sql.Open("sqlite3", config.SqlitePath)
	if err != nil {
		log.Printf("connectSqlite.Open: %v\n", err)
		return nil, err
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
	) WITHOUT ROWID;` // SQL Statement for Create Table

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
		Code:    "portsgrpc.Upsert-finished",
		Message: "Ports saved in the database.",
	}, nil
}

func main() {
	config := types.SetupConfig()
	log.Printf("main.SetupConfig: %#v\n", config)

	/***** Start three GreeterServers(with one of them to be the slowServer). *****/
	grpcAddress := strings.Split(config.GrpcAddress, ",")
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
