package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Encrypt aes 加密
func Encrypt(msg []byte, key string) (string, error) {

	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	bs := block.BlockSize()
	msg = PKCS7Padding(msg, bs)

	bm := cipher.NewCBCEncrypter(block, k[:bs])

	encrypted := make([]byte, len(msg))
	bm.CryptBlocks(encrypted, msg)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func Decrypt(encrypted string, key string) ([]byte, error) {

	crytedByte, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	k := []byte(key)

	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()

	bm := cipher.NewCBCDecrypter(block, k[:bs])

	orig := make([]byte, len(crytedByte))

	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	bm.CryptBlocks(orig, crytedByte)

	orig = PKCS7UnPadding(orig)
	return orig, nil
}

func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
