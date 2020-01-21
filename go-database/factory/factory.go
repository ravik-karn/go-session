package factory

import (
	"database/sql"
	"log"

	"go-database/config"
)

type Factory struct {
	db       *sql.DB
}

func New(conf *config.Config) *Factory {
	conn, err := sql.Open("postgres", conf.DBConnectionString)
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}
	defer conn.Close()
	return &Factory{db:conn}
}

func (f Factory) DB() *sql.DB {
	return f.db
}
