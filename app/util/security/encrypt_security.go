package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
)

func (s *Security) Encrypt(data interface{}) (string, error) {
	// Mengubah data menjadi JSON
	plainText, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	aesKey := []byte(s.encryptConfig.Aes)
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	return string(cipherText), nil
}

func (s *Security) Decrypt(cipherText string) ([]byte, error) {
	aesKey := []byte(s.encryptConfig.Aes)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("cipherText too short")
	}

	nonce, cipherTextByte := []byte(cipherText)[:nonceSize], []byte(cipherText)[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherTextByte, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
