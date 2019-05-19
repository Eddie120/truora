package api

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"log"
	"net/http"
	"truora/backend/app/helpers"
	"truora/backend/app/models"
	"truora/backend/config"
)

var db = config.GetConnetion()

func CreateKey(w http.ResponseWriter, r *http.Request) {

	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var modelKey models.Key
	error := decoder.Decode(&modelKey)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	publicKey := privateKey.PublicKey

	privateKeyInText := helpers.EncodePrivateKeyToString(privateKey)
	publicKeyInText := helpers.EncodePublicKeyToString(publicKey)

	privateKeyInAes256 := helpers.EncryptAES256([]byte(models.KEY),privateKeyInText)

	modelKey.PrivateKey = privateKeyInAes256
	modelKey.PublicKey = publicKeyInText

	query := "INSERT INTO m_keys (name,publickey,privatekey) VALUES ($1, $2, $3)"

	if _, error := db.Exec(query,modelKey.Name, modelKey.PrivateKey,modelKey.PrivateKey); error != nil {
		panic("it could not execute the next query "+ query +" : " + error.Error())
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(modelKey)
}


func Index(w http.ResponseWriter, r *http.Request) {

	defer db.Close()

	term := r.URL.Query().Get("text")

	 var query = "SELECT id, name FROM m_keys;"
	if term != "" {
		query = "SELECT id, name FROM m_keys WHERE lower(name) LIKE '%"+ term +"%'; "
	}

	rows, error := db.Query(query)

	if error != nil {
		panic("it could not execute the next query "+ query +" : " + error.Error())
	}

	defer rows.Close()

	var keys models.Keys
	for rows.Next() {

		var key models.Key

		error := rows.Scan(&key.Id,&key.Name)

		if  error != nil {
			panic("We can't scan the properties of key models : " + error.Error())
		}

		keys = append(keys, key)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(keys)
}


func Encrypt(w http.ResponseWriter, r *http.Request) {

	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var modelParams models.Params
	error := decoder.Decode(&modelParams)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	const query = "SELECT id,name,publickey FROM m_llaves WHERE id = $1;"

	rows, error := db.Query(query, modelParams.Id)

	if error != nil {
		panic("it could not execute the next : "+ query + " " + error.Error())
	}

	defer rows.Close()

	var key models.Key
	for rows.Next() {

		error := rows.Scan(&key.Id,&key.Name,&key.PublicKey)

		if  error != nil {
			panic("We can't scan the properties of key models : " + error.Error())
		}
	}

	if key.PublicKey != "" {

		block, _ := pem.Decode([]byte(key.PublicKey))
		publicKey, error := x509.ParsePKCS1PublicKey(block.Bytes)

		if error != nil {
			panic("it could not analize the public key" + error.Error())
		}

		messageEncrypt := helpers.Encrypt(modelParams.Text,publicKey);

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(messageEncrypt)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func Decrypt(w http.ResponseWriter, r *http.Request) {

	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	var modelParams models.Params
	error := decoder.Decode(&modelParams)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const query = "SELECT id,name,privatekey FROM m_keys WHERE id = $1;"

	rows, error := db.Query(query, modelParams.Id)

	if error != nil {
		panic("it could not execute the next query "+ query + " " + error.Error())
	}

	defer rows.Close()

	var key models.Key
	for rows.Next() {

		error := rows.Scan(&key.Id,&key.Name,&key.PrivateKey)

		if  error != nil {
			panic("it could not scan properties of key model " + error.Error())
		}
	}

	if key.PrivateKey != "" {

		privateKeyWithOutAes256 := helpers.DecryptAES256([]byte(models.KEY),key.PrivateKey)

		block, _ := pem.Decode([]byte(privateKeyWithOutAes256))
		privateKey, error := x509.ParsePKCS1PrivateKey(block.Bytes)

		if error != nil {
			log.Fatal("We have some problems loading privateKey ",error.Error())
		}

		message := helpers.Decrypt(modelParams.Text,privateKey);

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

