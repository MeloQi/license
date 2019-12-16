package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"bytes"
	"strings"
)

func GenRsaKey(bits int) (pubkey, prikey string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prikeyPem := &bytes.Buffer{}
	err = pem.Encode(prikeyPem, block)
	if err != nil {
		return "", "", err
	}

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkeyPem := &bytes.Buffer{}
	err = pem.Encode(pubkeyPem, block)
	if err != nil {
		return "", "", err
	}
	return string(pubkeyPem.Bytes()), string(prikeyPem.Bytes()), nil
}

func GetStrByPem(KeyPem string) string {
	keyArry := strings.Split(KeyPem, "\n")
	var KeyStr string
	for i := 1; i < len(keyArry)-2; i++ {
		KeyStr += keyArry[i]
	}
	return KeyStr
}

func GetPemByStr(KeyStr string) string {
	var KeyPem string
	KeyPem += "-----BEGIN PUBLIC KEY-----\n"
	var i int = 0
	for ; i < len(KeyStr)/64; i++ {
		KeyPem += KeyStr[i*64:i*64+64] + "\n"
	}
	KeyPem += KeyStr[i*64:] + "\n"
	KeyPem += "-----END PUBLIC KEY-----\n"

	return KeyPem
}
