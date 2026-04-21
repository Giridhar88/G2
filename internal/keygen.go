package internal

import (
	"golang.org/x/crypto/scrypt"
)

// Genearte the key depending on the salt in the cs struct and the pwd
func GenerateKey(pwd []byte, salt []byte) ([]byte, error) {
	genkey, err := scrypt.Key(pwd, salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	return genkey, nil
}
