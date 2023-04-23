package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"go_cases/tools/cryptor"
)

type cbc struct {
	key []byte
	iv  []byte
}

func NewCbc(key string) cryptor.Cryptor {
	return &cbc{key: []byte(key)}
}

func (c *cbc) EncryptString(plainText string) (string, error) {
	cipherText, err := c.Encrypt([]byte(plainText))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (c *cbc) Encrypt(plainText []byte) ([]byte, error) {
	if len(c.key) != 16 && len(c.key) != 24 && len(c.key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	paddingText := pkcs7Padding([]byte(plainText), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, c.key[:blockSize])
	ciphertext := make([]byte, len(paddingText))
	blockMode.CryptBlocks(ciphertext, paddingText)

	return ciphertext, nil
}

func (c *cbc) DecryptString(ciphertext string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	plainText, err := c.Decrypt(cipherBytes)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func (c *cbc) Decrypt(cipherText []byte) ([]byte, error) {
	if len(c.key) != 16 && len(c.key) != 24 && len(c.key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, c.key[:blockSize])
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)

	return pkcs7UnPadding(paddingText), nil
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	var paddingText []byte
	if padding == 0 {
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}

	return append(ciphertext, paddingText...)
}

func pkcs7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unPadding := int(plainText[length-1])
	return plainText[:(length - unPadding)]
}

func paddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}
