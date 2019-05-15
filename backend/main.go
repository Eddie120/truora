package main

import (
	"log"
	"net/http"
	"truora/backend/routes"
)

func main() {

	r := routes.CargarRutas()

	server := http.ListenAndServe(":3000", r)

	log.Fatal(server)
}