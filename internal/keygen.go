package internal

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

// struct holds the key generated when the main function runs
type crypto struct {
	pwd  []byte
	salt []byte
}

// helper struct used in encrypt and decrypt function it holds the salt and the key, salt is of 16 bytes
var cs *crypto

// initializes the struct with the password and a random salt available in the internal package
func Init(bp []byte) {
	rsalt := make([]byte, 16)
	rand.Read(rsalt)
	cs = &crypto{
		pwd:  bp,
		salt: rsalt,
	}
}

// Genearte the key depending on the salt in the cs struct and the pwd
func (cs *crypto) GenerateKey() ([]byte, error) {
	genkey, err := scrypt.Key(cs.pwd, cs.salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	return genkey, nil
}
