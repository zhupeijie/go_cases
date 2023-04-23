package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"go_cases/tools/cryptor"
)

type ctr struct {
	key []byte
	iv  []byte
}

func NewCtr(key string) cryptor.Cryptor {
	return &ctr{key: []byte(key)}
}

func (c *ctr) EncryptString(plainText string) (string, error) {
	cipherText, err := c.Encrypt([]byte(plainText))
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(cipherText), nil
}

func (c *ctr) Encrypt(plainText []byte) ([]byte, error) {
	if len(c.key) != 16 && len(c.key) != 24 && len(c.key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plainBytes := []byte(plainText)
	stream := cipher.NewCTR(block, c.key[:blockSize])
	cipherText := make([]byte, len(plainBytes))
	stream.XORKeyStream(cipherText, plainBytes)

	return cipherText, nil
}

func (c *ctr) DecryptString(ciphertext string) (string, error) {
	cipherBytes, err := base64.RawURLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	plainText, err := c.Decrypt(cipherBytes)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func (c *ctr) Decrypt(cipherText []byte) ([]byte, error) {
	if len(c.key) != 16 && len(c.key) != 24 && len(c.key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	stream := cipher.NewCTR(block, c.key[:blockSize])
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}
