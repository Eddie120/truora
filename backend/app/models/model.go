package models

import "net/http"

const KEY = "AES256Key-32Characters1234567890"

type Llave struct {
	Id string `json:"id"`
	Nombre string `json:"nombre"`
	LlavePrivada string `json:"llaveprivada"`
	LlavePublica string `json:"llavepublica"`
}

type Parametro struct {
	Id string `json:"id"`
	Texto string `json:"texto"`
}

type Llaves []Llave


type Ruta struct {
	Patron  string
	Metodo  string
	Funcion http.HandlerFunc
}

type Rutas []Ruta