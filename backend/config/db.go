package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)


func GetConnection() (db *sql.DB) {

	db, err := sql.Open("postgres",  "postgresql://eddie@localhost:26257/bank?sslmode=disable")

	if err != nil {
		panic("it could not connect with database " + err.Error())
	}

	return db
}


