package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// decrypt the bytes given to the func
func Decrypt(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return []byte{}, nil
	}
	cs.salt = data[:16]
	key, err := cs.GenerateKey()
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()

	nonce := data[16 : 16+nonceSize]
	cipherText := data[16+nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		fmt.Println("failed to decrypt the contents")
		return nil, err
	}

	return plainText, nil
}
