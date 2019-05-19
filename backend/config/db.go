package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)


func GetConnetion() (db *sql.DB) {
	//db, error := sql.Open("postgres",  "postgresql://eddie@localhost:26257/bank?ssl=false&sslmode=disable&sslrootcert=certs/ca.crt&sslkey=certs/client.eddie.key&sslcert=certs/client.eddie.crt")
	// En local tenia que conectarme de manera insegura el nodo ya que tengo problemas con un password
	db, error := sql.Open("postgres",  "postgresql://eddie@localhost:26257/bank?sslmode=disable")

	if error != nil {
		log.Fatal("Se presentó un error durante la conexión a la base de datos: ", error)
	}

	return db
}


