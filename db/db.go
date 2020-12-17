package db

import (
	"database/sql"
	"log"
	"../types"
	model "./types"
)

type DBConn struct {
	DB *sql.DB
}

//Connect - SQLite connection routine
func Connect(_config types.Config) (*DBConn, error) {
	db, err := sql.Open("sqlite3", _config.SQLitePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	//Creating table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS ports (
		port_id char(5) PRIMARY KEY,
		name varchar(100) NOT NULL,
		country varchar(100) NOT NULL,
		alias varchar(500),
		regions varchar(500),
		coordinates varchar(100),
		province varchar(100),
		timezone varchar(50),
		unlocs varchar(500),
		code varchar (100)
	) WITHOUT ROWID;
	CREATE INDEX [idx_ports_name] ON "ports" ([port_id]);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}

	return &DBConn{DB: db}, nil
}

func (db *DBConn) Exists(PortId string) (bool, error){
	stmt, err := db.DB.Prepare("select 1 from ports where port_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var exists int
	err = stmt.QueryRow(PortId).Scan(&exists)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return exists != 0, nil
}

//Upset insert or update record
func (db *DBConn) Upset(body model.Ports) (bool, error) {
	exists, err := db.Exists(body.PortId)
	if err != nil {
		return false, err
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}
	if !exists {
		stmt, err := tx.Prepare(`insert into ports 
		(port_id,name,country,alias,regions,coordinate,province,timezone,unlocs,code) 
		values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(body.PortId, body.Name, body.Alias, body.Regions, body.Coordinates, body.Province, body.Timezone, body.Unlocs, body.Code)
		if err != nil {
			return false, err
		}
	} else {
		stmt, err := tx.Prepare(`update ports set 
		name=?,country=?,alias=?,regions=?,coordinate=?,province=?,timezone=?,unlocs=?,code=? 
		where port_id=?`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(body.Name, body.Alias, body.Regions, body.Coordinates, body.Province, body.Timezone, body.Unlocs, body.Code, body.PortId)
		if err != nil {
			return false, err
		}
	}

	tx.Commit()

	return true, nil
}