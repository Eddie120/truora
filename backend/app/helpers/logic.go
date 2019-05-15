package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

func EncodePrivateKeyToString(key *rsa.PrivateKey) string {

	err := key.Validate()

	if err != nil {
		fmt.Println("Error validando la llave privada: %s\n", err)
	}

	 privateKey := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	privateKeyString := string(pem.EncodeToMemory(privateKey))

	return privateKeyString
}

func EncodePublicKeyToString(pubkey rsa.PublicKey) string {
	publicKeyBytes, _ := asn1.Marshal(pubkey)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicKeyString := string(pem.EncodeToMemory(pemkey))

	return publicKeyString
}

/***
 Para ver si las llaves de generaban de manera correcta
 */
/*func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, _ := os.Create(fileName)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	_ = pem.Encode(outFile, privateKey)
}*/

func Encrypt(mensaje string, llavePublica *rsa.PublicKey) (string) {
	mensajeCifrado, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, llavePublica, []byte(mensaje),nil)
	if err != nil {
		fmt.Println("Error encriptando el mensaje con la llave publica : %s\n", err)
	}
	return base64.StdEncoding.EncodeToString(mensajeCifrado)
}

func Decrypt(mensajeCifrado string, llavePrivada *rsa.PrivateKey) (string) {
	mensaje,err := base64.StdEncoding.DecodeString(mensajeCifrado)
	if err != nil {
		fmt.Println("Error decodificando string: %s\n", err)
	}

	mensajeDescifrado, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, llavePrivada, mensaje,nil)

	if err != nil {
		fmt.Println("Error desencriptando el mensaje el mensaje: %s\n", err)
	}

	return string(mensajeDescifrado)
}

func EncryptAES256(key []byte, privateKey string) string {

	privateKeyByte := []byte(privateKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creando un nuevo cipher: %s\n", err)
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(privateKeyByte))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], privateKeyByte)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func DecryptAES256(key []byte, privateKey string) string {
	cipherPrivateKey, err := base64.URLEncoding.DecodeString(privateKey)

	if err != nil {
		fmt.Println("Error decodificando la llave privada: %s\n", err)
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(cipherPrivateKey) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		panic(err)
	}

	iv := cipherPrivateKey[:aes.BlockSize]
	cipherPrivateKey = cipherPrivateKey[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherPrivateKey, cipherPrivateKey)

	privateKey = string(cipherPrivateKey)

	return privateKey
}
