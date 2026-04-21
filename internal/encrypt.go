package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// encrypt the bytes given to the func
func Encrypt(data []byte, pwd []byte) ([]byte, error) {
	if len(data) == 0 {
		return []byte{}, nil
	}
	salt := make([]byte, 16)
	rand.Read(salt)
	key, err := GenerateKey(pwd, salt)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)
	ciphertext := gcm.Seal(nil, nonce, data, nil)
	out := append(salt, nonce...)
	out = append(out, ciphertext...)

	return out, nil
}
