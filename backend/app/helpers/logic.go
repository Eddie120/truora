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
	"io"
)

func EncodePrivateKeyToString(key *rsa.PrivateKey) string {

	error := key.Validate()

	if error != nil {
		panic("We can not validate private key "+ error.Error())
	}

	 privateKey := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	privateKeyString := string(pem.EncodeToMemory(privateKey))

	return privateKeyString
}

func EncodePublicKeyToString(pubkey rsa.PublicKey) string {
	publicKeyBytes, error := asn1.Marshal(pubkey)

	if error != nil {
		panic("We can not transform public key to bytes " + error.Error())
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicKeyString := string(pem.EncodeToMemory(pemkey))

	return publicKeyString
}


func Encrypt(message string, publicKey *rsa.PublicKey) (string) {
	messageEncrypted, error := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(message),nil)

	if error != nil {
		panic("We can not encrypt message with the public key " + error.Error())
	}

	return base64.StdEncoding.EncodeToString(messageEncrypted)
}

func Decrypt(messageEncrypt string, privateKey *rsa.PrivateKey) (string) {
	messageEncrypted, error := base64.StdEncoding.DecodeString(messageEncrypt)

	if error != nil {
		panic("We can not decode the message with base64 " + error.Error())
	}

	messageDecrypted, error := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, messageEncrypted,nil)

	if error != nil {
		panic("We can not decrypt the message " + error.Error())
	}

	return string(messageDecrypted)
}

func EncryptAES256(key []byte, privateKey string) string {

	privateKeyByte := []byte(privateKey)

	block, error := aes.NewCipher(key)
	if error != nil {
		panic("it could not create a block to the key " + string(key) + " " + error.Error())
	}

	ciphertext := make([]byte, aes.BlockSize+len(privateKeyByte))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("We could not read all bytes "+ err.Error())
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], privateKeyByte)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func DecryptAES256(key []byte, privateKey string) string {
	cipherPrivateKey, error := base64.URLEncoding.DecodeString(privateKey)

	if error != nil {
		panic("We could not read all bytes "+ error.Error())
	}

	block, error := aes.NewCipher(key)
	if error != nil {
		panic("it could not create a block to the key " + string(key) + " " + error.Error())
	}

	if len(cipherPrivateKey) < aes.BlockSize {
		error = errors.New("the block size is really small")
		panic(error)
	}

	iv := cipherPrivateKey[:aes.BlockSize]
	cipherPrivateKey = cipherPrivateKey[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherPrivateKey, cipherPrivateKey)

	privateKey = string(cipherPrivateKey)

	return privateKey
}
