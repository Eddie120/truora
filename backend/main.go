package main

import (
	"log"
	"net/http"
	"truora/backend/routes"
)

func main() {

	r := routes.LoadRoutes()

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatal(err)
	}

}