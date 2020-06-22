package pki

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func GetPrivateKey() (*rsa.PrivateKey, error) {
	// Read the private key
	key, err := ioutil.ReadFile("./certs/pri.pem")
	if err != nil {
		return nil, err
	}
	// Extract the PEM-encoded data block
	block, _ := pem.Decode(key)
	if block == nil {
		log.Fatalf("bad key data: %s\n", "not PEM-encoded")
	}
	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		log.Fatalf("unknown key type %q, want %q", got, want)
	}
	der, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("bad private key: %s", err)
		return nil, err
	}
	return der, err
}
