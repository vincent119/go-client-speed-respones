package token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	//"encoding/base64"
	"encoding/pem"
	"errors"
	//"fmt"
)

//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDBuX8h/lJAFeatlhVZqCRKLnAlApVS+FSAuII89bvy3J6V1WM5
019z36fQe/+PyBtKT2bMIcUJoJklg1i9yCCi2g6DjngUFUL1o30Q+5qnF+4ZDRsZ
MZc0pJeL1RrbPjYgudzMgqZxyXyxLpx9t282nVU3+2MMR01cby4fZiroYwIDAQAB
AoGAdy2lFsac1ywiE9xnGhck7DYRa9NhBgAcDGx2QwZyMACl10vFlYeCC4kmdk4y
dMQIdfjPJvoHe9q/HKLnbym1WSh9MpCoPh9/27zUg7NRQoqeHXGT3gCNvHRBAxwL
L2PBWxddxq4Mic88MGh9jWSiwd+sA62PJ6bIAEhCppAzqlECQQDigr6mRDU50M87
IOsfuJ28hKbc6+uPMxVT+fUBXPiwimm80TMv1ajoJrGpxePPYXKr3ax4/BMvlVx/
Y57fNconAkEA2vIKYmm/Y/+yyowfGqXyEVwPlQThYvlHZKU2b0caCtCp5CMRA1mU
ZYHwZxjt1U8ch25A3fuBi1VCCJYZv5gBZQJBALtruXpzx6K+Hypqs4I8zO+Wx0bX
QsLrXsNseIiXOANDk/gzFWqIHMlycuDqJ/SJSkvhEPvOf8WIjc+uS1+f71MCQQDN
6OIxSeiwoDLC4OMDyduNzfntmShrT6uAEQzPvJvWXgXZQ81lTCMPEBRKsZxDh2J6
UAt7eWSM6ILw1lGi6c8tAkBklNOgX+RzfSXOuNcHXGha78vwGSzKOfMliYKwOB/B
XSwSv4EiEJvAJ3ndE73omY1JnVam8mm9yrYbmXqGwKbo
-----END RSA PRIVATE KEY-----`)

//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDBuX8h/lJAFeatlhVZqCRKLnAl
ApVS+FSAuII89bvy3J6V1WM5019z36fQe/+PyBtKT2bMIcUJoJklg1i9yCCi2g6D
jngUFUL1o30Q+5qnF+4ZDRsZMZc0pJeL1RrbPjYgudzMgqZxyXyxLpx9t282nVU3
+2MMR01cby4fZiroYwIDAQAB
-----END PUBLIC KEY-----`)

func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

/*
func example () {
	data, _ := RsaEncrypt([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, _ := RsaDecrypt(data)
	fmt.Println(string(origData))
} */
