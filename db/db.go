package db

import (
	"../types"
	model "./types"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var config *types.Config
type Conn struct {
	DB *sql.DB
}

//Connect - SQLite connection routine
func Connect() (*Conn, error) {
	if config == nil {
		err := errors.New("Database configuration not informed: config is null")
		log.Println(err)
		return nil, err
	}

	if _, err := os.Stat(config.SQLitePath); err != nil {
		log.Printf("%v - creating new database file\n", err)
		file, err := os.Create(config.SQLitePath) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Printf("Database file '%s' created", config.SQLitePath)
	}

	db, err := sql.Open("sqlite3", config.SQLitePath)
	if err != nil {
		log.Printf("db.Connect: %v\n", err)
		return nil, err
	}
	defer db.Close()
	log.Printf("db.Connect-db: %#v\n", db)

	// create table if not exists
	createTable(db)

	return &Conn{DB: db}, nil
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS ports (
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
	) WITHOUT ROWID;` // SQL Statement for Create Table

	log.Println("Create 'ports' table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Printf("db.createTable", err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("'ports' table created")
}

func (c *Conn) Exists(PortId string) (bool, error){
	log.Printf("db.Exists-db: %#v\n", c.DB)

	row, err := c.DB.Query("SELECT port_id FROM ports WHERE port_id = ?", PortId)
	if err != nil {
		log.Printf("db.Exists: %v\n", err)
		return false, err
	}
	defer row.Close()

	return row.Next(), nil
}

//Upset insert or update record
func (c *Conn) Upset(body model.Ports) (bool, error) {
	//exists, err := c.Exists(body.PortId)
	//if err != nil {
	//	log.Printf("db.Upset: %v\n", err)
	//	return false, err
	//}
	exists := false

	tx, err := c.DB.Begin()
	if err != nil {
		log.Printf("db.Upset: %v\n", err)
		return false, err
	}
	if !exists {
		stmt, err := tx.Prepare(`insert into ports 
		(port_id,name,country,alias,regions,coordinates,province,timezone,unlocs,code) 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {
			log.Printf("db.Upset: %v\n", err)
			return false, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(body.PortId, body.Name, body.Alias, body.Regions, body.Coordinates, body.Province, body.Timezone, body.Unlocs, body.Code)
		if err != nil {
			log.Printf("db.Upset: %v\n", err)
			return false, err
		}
	} else {
		stmt, err := tx.Prepare(`update ports set 
		name=?,country=?,alias=?,regions=?,coordinates=?,province=?,timezone=?,unlocs=?,code=? 
		where port_id=?`)
		if err != nil {
			log.Printf("db.Upset: %v\n", err)
			return false, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(body.Name, body.Alias, body.Regions, body.Coordinates, body.Province, body.Timezone, body.Unlocs, body.Code, body.PortId)
		if err != nil {
			log.Printf("db.Upset: %v\n", err)
			return false, err
		}
	}

	if err = tx.Commit(); err != nil {
		log.Printf("db.Upset: %v\n", err)
		return false, err
	}

	return true, nil
}

func SetConfig(_config *types.Config) {
	config = _config
}