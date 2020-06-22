package pki

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"log"
	"fmt"
)

func Encrypt(message string, priv *rsa.PrivateKey) string {
	inp := []byte(message)
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, &priv.PublicKey, inp)
	if err != nil {
		log.Fatalf("encrypt Error: %s\n", err)
	}
	fmt.Printf("Encrypted: %s \n", base64.StdEncoding.EncodeToString(encrypted))
	return base64.StdEncoding.EncodeToString(encrypted)
}

func Decrypt(message string, priv *rsa.PrivateKey) []byte {
	encrypted, err := base64.StdEncoding.DecodeString(message)
	// Decrypt the data
	out, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encrypted)
	if err != nil {
		log.Fatalf("decrypt Error: %s\n", err)
	}
	return out
}