package database

import (
	"bytes"
	"fmt"

	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Dialect  string
	Host     string
	Name     string
	Username string
	Password string
	Port     string
}

type Database struct {
	*sqlx.DB
}

func LoadDatabase(config DatabaseConfig) (database *Database, err error) {
	var buffer bytes.Buffer
	buffer.WriteString(config.Dialect + "://")
	buffer.WriteString(config.Username + ":" + config.Password)
	buffer.WriteString("@")
	buffer.WriteString(config.Host + ":" + config.Port + "/")
	buffer.WriteString(config.Name)
	buffer.WriteString("?sslmode=require")
	connectionString := buffer.String()

	db, err := sqlx.Open(config.Dialect, connectionString)
	if err != nil {
		err = fmt.Errorf("failed to connect to database. %s", err.Error())
		return
	}

	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
		// return
	}

	database = &Database{
		db,
	}

	return
}
