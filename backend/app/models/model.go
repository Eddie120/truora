package models

import "net/http"

const KEY = "AES256Key-32Characters1234567890"

type Key struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type Params struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type Keys []Key

type Route struct {
	Pattern  string
	Method   string
	Function http.HandlerFunc
}

type Routes []Route
