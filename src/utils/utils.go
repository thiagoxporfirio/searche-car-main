package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

func CreateCodeId() string {
	return uuid.NewV4().String()
}

func Encrypt(code string) string {
	key := []byte("CarsSearch2022")
	plaintext := []byte(code)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("CS")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	ciphertextStr := string(ciphertext[:])
	return ciphertextStr
}

func decrypt(code string) string {
	key := []byte("CarsSearch2022")
	ciphertext, _ := hex.DecodeString(code)
	nonce := []byte("CS")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	plaintextStr := string(plaintext[:])
	return plaintextStr
}
