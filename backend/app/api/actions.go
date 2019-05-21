package api

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"truora/backend/app/helpers"
	"truora/backend/app/models"
	"truora/backend/config"
)


var (
	db = config.GetConnection()
	err error
	rows *sql.Rows
	publicKey *rsa.PublicKey
	privateKey *rsa.PrivateKey
)


func CreateKey(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var modelKey models.Key
	err = decoder.Decode(&modelKey)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	publicKey := privateKey.PublicKey

	privateKeyInText := helpers.EncodePrivateKeyToString(privateKey)
	publicKeyInText := helpers.EncodePublicKeyToString(publicKey)

	privateKeyInAes256 := helpers.EncryptAES256([]byte(models.KEY), privateKeyInText)

	modelKey.PrivateKey = privateKeyInAes256
	modelKey.PublicKey = publicKeyInText

	query := "INSERT INTO m_keys (name,publickey,privatekey) VALUES ($1, $2, $3)"

	if _, err = db.Exec(query, modelKey.Name, modelKey.PublicKey, modelKey.PrivateKey); err != nil {
		panic("it could not execute the next query " + query + " : " + err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(modelKey)
}

func Index(w http.ResponseWriter, r *http.Request) {

	term := r.URL.Query().Get("text")
	perPage := r.URL.Query().Get("perPage")
	//lastId := chi.URLParam(r, "lastId")

	query := "SELECT id, name FROM m_keys ORDER BY id ASC LIMIT $1"
	rows, err =  db.Query(query, perPage)

	if term != "" {
		query = "SELECT id, name FROM m_keys WHERE lower(name) LIKE '%' || $1 || '%' ORDER BY id ASC LIMIT $2 "
		rows, err =  db.Query(query, term, perPage)
	}

	defer rows.Close()

	if err != nil {
		panic("it could not execute the next query " + query + " : " + err.Error())
	}

	var keys models.Keys
	for rows.Next() {

		var key models.Key

		err = rows.Scan(&key.Id, &key.Name)

		if err != nil {
			panic("We can't scan the properties of key models : " + err.Error())
		}

		keys = append(keys, key)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(keys)
}

func Encrypt(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var modelParams models.Params
	err = decoder.Decode(&modelParams)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := "SELECT id,name,publickey FROM m_keys WHERE id = $1;"

	var key models.Key
	err = db.QueryRow(query, modelParams.Id).Scan(&key.Id, &key.Name, &key.PublicKey)

	if err != nil {
		panic("it could not execute the next : " + query + " " + err.Error())
	}


	if key.PublicKey != "" {

		block, _ := pem.Decode([]byte(key.PublicKey))
		publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)

		if err != nil {
			panic("it could not analize the public key" + err.Error())
		}

		messageEncrypt := helpers.Encrypt(modelParams.Text, publicKey)

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(messageEncrypt)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func Decrypt(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	var modelParams models.Params
	err = decoder.Decode(&modelParams)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := "SELECT id,name,privatekey FROM m_keys WHERE id = $1;"

	var key models.Key
	err = db.QueryRow(query, modelParams.Id).Scan(&key.Id, &key.Name, &key.PrivateKey)

	if err != nil {
		panic("it could not execute the next query " + query + " " + err.Error())
	}

	if key.PrivateKey != "" {

		privateKeyWithOutAes256 := helpers.DecryptAES256([]byte(models.KEY), key.PrivateKey)

		block, _ := pem.Decode([]byte(privateKeyWithOutAes256))
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)

		if err != nil {
			panic("it could not load the private Key " + err.Error())
		}

		message := helpers.Decrypt(modelParams.Text, privateKey)

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(message)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
