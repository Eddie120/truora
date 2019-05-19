package main

import (
	"net/http"
	"truora/backend/routes"
)

func main() {

	r := routes.LoadRoutes()

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		panic("it could not listen TCP network " + err.Error())
	}

}