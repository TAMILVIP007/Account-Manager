package src

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func encryptAES(plaintext string) ([]byte, error) {
	encyptkey := []byte(Envars.Encyptkey)
	block, err := aes.NewCipher(encyptkey)
	if err != nil {
		return nil, err
	}
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	plaintext += fmt.Sprintf("%c", padding)
	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, []byte(plaintext))
	return ciphertext, nil
}

func decryptAES(ciphertext []byte) (string, error) {
	encyptkey := []byte(Envars.Encyptkey)
	block, err := aes.NewCipher(encyptkey)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
