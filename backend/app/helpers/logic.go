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

	err := key.Validate()

	if err != nil {
		panic("We can not validate private key "+ err.Error())
	}

	 privateKey := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	privateKeyString := string(pem.EncodeToMemory(privateKey))

	return privateKeyString
}

func EncodePublicKeyToString(pubkey rsa.PublicKey) string {
	publicKeyBytes, err := asn1.Marshal(pubkey)

	if err != nil {
		panic("We can not transform public key to bytes " + err.Error())
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicKeyString := string(pem.EncodeToMemory(pemkey))

	return publicKeyString
}


func Encrypt(message string, publicKey *rsa.PublicKey) (string) {
	messageEncrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(message),nil)

	if err != nil {
		panic("We can not encrypt message with the public key " + err.Error())
	}

	return base64.StdEncoding.EncodeToString(messageEncrypted)
}

func Decrypt(messageEncrypt string, privateKey *rsa.PrivateKey) (string) {
	messageEncrypted, err := base64.StdEncoding.DecodeString(messageEncrypt)

	if err != nil {
		panic("We can not decode the message with base64 " + err.Error())
	}

	messageDecrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, messageEncrypted,nil)

	if err != nil {
		panic("We can not decrypt the message " + err.Error())
	}

	return string(messageDecrypted)
}

func EncryptAES256(key []byte, privateKey string) string {

	privateKeyByte := []byte(privateKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("it could not create a block to the key " + string(key) + " " + err.Error())
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
	cipherPrivateKey, err := base64.URLEncoding.DecodeString(privateKey)

	if err != nil {
		panic("We could not read all bytes "+ err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("it could not create a block to the key " + string(key) + " " + err.Error())
	}

	if len(cipherPrivateKey) < aes.BlockSize {
		err = errors.New("the block size is really small")
		panic(err)
	}

	iv := cipherPrivateKey[:aes.BlockSize]
	cipherPrivateKey = cipherPrivateKey[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherPrivateKey, cipherPrivateKey)

	privateKey = string(cipherPrivateKey)

	return privateKey
}
