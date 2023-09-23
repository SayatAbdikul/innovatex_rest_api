package server

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "postgres://sayat:Sayat2005!@localhost:5432/innovatex?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err := DB.Ping(); err != nil {
		panic(err)
	}

}
