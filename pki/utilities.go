package pki

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func GenRsaPair() *rsa.PrivateKey {

	reader := rand.Reader
	bitSize := 512

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		fmt.Println("Can't generate rsa key pair?")
	}
	return key
}
