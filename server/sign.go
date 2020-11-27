package server

import (
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	mrand "math/rand"
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charSet[mrand.Intn(len(charSet))]
	}
	return string(b)
}

func sign(msg string) string {
	key, _ := ioutil.ReadFile("serverKey.pem")
	block, _ := pem.Decode(key)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	rng := crand.Reader
	hashed := sha256.Sum256([]byte(msg))
	signature, _ := rsa.SignPKCS1v15(rng, privateKey, crypto.SHA256, hashed[:])
	return base64.StdEncoding.EncodeToString(signature)
}
