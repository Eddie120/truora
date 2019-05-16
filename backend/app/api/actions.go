package api

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"truora/backend/config"
	"truora/backend/app/models"
	"truora/backend/app/helpers"
)

func Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Conexion realizada con exito"))
}

/*
	Crea una llave
	Cifrando la llave privada en AES-256
 */
func CrearLlave(w http.ResponseWriter, r *http.Request) {
	db := config.GetConexion()

	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var modeloLlave models.Llave
	error := decoder.Decode(&modeloLlave)

	if error != nil {
		panic(error)
	}

	defer r.Body.Close()

	llavePrivada, _ := rsa.GenerateKey(rand.Reader, 4096)
	llavePublica := llavePrivada.PublicKey

	llavePrivadaEnTexto := helpers.EncodePrivateKeyToString(llavePrivada)
	llavePublicaEnTexto := helpers.EncodePublicKeyToString(llavePublica)

	fmt.Println("Llave privada generada \n",llavePrivadaEnTexto)
	fmt.Println("Llave publica generada \n",llavePublicaEnTexto)

	llavePrivadaAES256 := helpers.EncryptAES256([]byte(models.KEY),llavePrivadaEnTexto)
	fmt.Println("Llave privada AES-256 \n",llavePrivadaAES256)

	llavePrivadaSinAES256 := helpers.DecryptAES256([]byte(models.KEY),llavePrivadaAES256)
	fmt.Println("Llave privada sin AES-256 \n",llavePrivadaSinAES256)


	modeloLlave.LlavePrivada = llavePrivadaAES256
	modeloLlave.LlavePublica = llavePublicaEnTexto

	//creamos la tabla si no existe
	 _, err := db.Exec("CREATE TABLE IF NOT EXISTS m_Llaves(id serial PRIMARY KEY,nombre VARCHAR(255) NOT NULL,llavepublica TEXT NOT NULL,llaveprivada TEXT NOT NULL)")

	 if err != nil {
		 log.Fatal(err)
	 }

	if _, err := db.Exec(
		"INSERT INTO m_Llaves (nombre,llavepublica,llaveprivada) VALUES ($1, $2, $3)",modeloLlave.Nombre, modeloLlave.LlavePublica,modeloLlave.LlavePrivada); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated) // 202
	json.NewEncoder(w)
}


/*
	Permite listar
	todas las llaves disponibles
 */
func ListarLlaves(w http.ResponseWriter, r *http.Request) {
	db := config.GetConexion()

	defer db.Close()

	term := r.URL.Query().Get("texto")

	if term != "" {
		fmt.Println("Este el termino que estas buscando", term)
	}

	 var query = "SELECT id, nombre FROM m_Llaves"
	if term != "" {
		query = "SELECT id, nombre FROM m_Llaves WHERE lower(nombre) LIKE '%"+ term +"%'; "
	}

	//fmt.Println(query)
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var llaves models.Llaves
	for rows.Next() {

		var llave models.Llave

		err := rows.Scan(&llave.Id,&llave.Nombre)

		if  err != nil {
			log.Fatal(err)
		}

		llaves = append(llaves, llave)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(llaves)
}

/*
	Encripta el mensaje con
	la llave publica
 */
func EncriptarMensaje(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var modeloParametros models.Parametro
	error := decoder.Decode(&modeloParametros)

	if error != nil {
		panic(error)
	}

	db := config.GetConexion()

	defer db.Close()

	const query = `SELECT id,nombre,llavepublica FROM m_llaves WHERE id = $1;`

	rows, error := db.Query(query, modeloParametros.Id)

	if error != nil {
		panic(error)
	}

	defer rows.Close()

	var llave models.Llave
	for rows.Next() {

		err := rows.Scan(&llave.Id,&llave.Nombre,&llave.LlavePublica)

		if  err != nil {
			log.Fatal(err)
		}
	}

	//fmt.Println("Llave \n",llave)

	if llave.LlavePublica != "" {

		block, _ := pem.Decode([]byte(llave.LlavePublica))
		llavePublica, _ := x509.ParsePKCS1PublicKey(block.Bytes)

		mensajeEncriptado := helpers.Encrypt(modeloParametros.Texto,llavePublica);

		//fmt.Println("Llave publica generada \n",llavePublica)
		//fmt.Println("Mensaje cifrado generado \n",mensajeEncriptado)

		w.WriteHeader(http.StatusOK) // 202
		json.NewEncoder(w).Encode(mensajeEncriptado)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		json.NewEncoder(w)
	}
}

func DesencriptarMensaje(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var modeloParametros models.Parametro
	error := decoder.Decode(&modeloParametros)

	if error != nil {
		panic(error)
	}

	//fmt.Println(modeloParametros)
	fmt.Println(modeloParametros)

	db := config.GetConexion()

	defer db.Close()

	const query = `SELECT id,nombre,llaveprivada FROM m_llaves WHERE id = $1;`

	rows, error := db.Query(query, modeloParametros.Id)

	if error != nil {
		panic(error)
	}

	defer rows.Close()

	var llave models.Llave
	for rows.Next() {

		err := rows.Scan(&llave.Id,&llave.Nombre,&llave.LlavePrivada)

		if  err != nil {
			log.Fatal(err)
		}
	}

	//fmt.Println("Llave privada  es  \n",llave.LlavePrivada)

	if llave.LlavePrivada != "" {

		llavePrivadaSinAES256 := helpers.DecryptAES256([]byte(models.KEY),llave.LlavePrivada)

		block, _ := pem.Decode([]byte(llavePrivadaSinAES256))
		llavePrivada, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
		mensajeOriginal := helpers.Decrypt(modeloParametros.Texto,llavePrivada);

		//fmt.Println("Llave privada con AES-256 \n",llave.LlavePrivada)
		//fmt.Println("Llave privada sin AES-256 \n",llavePrivadaSinAES256)
		//fmt.Println("Mensaje original \n",reflect.TypeOf(mensajeOriginal))
		fmt.Println("Mensaje original \n",mensajeOriginal)

		//w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK) // 202
		json.NewEncoder(w).Encode(mensajeOriginal)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		json.NewEncoder(w)
	}
}

