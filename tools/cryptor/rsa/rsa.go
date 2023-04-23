package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"runtime"
)

type encryptor struct {
	pubKey []byte
}

type decryptor struct {
	priKey []byte
}

func NewEncryptor(publicKey string) *encryptor {
	return &encryptor{
		pubKey: []byte(publicKey),
	}
}

func NewDecryptor(privateKey string) *decryptor {
	return &decryptor{
		priKey: []byte(privateKey),
	}
}

func (r *encryptor) EncryptString(plainText string) (string, error) {
	cipherText, err := r.Encrypt([]byte(plainText))
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(cipherText), nil
}

func (r *encryptor) Encrypt(plainText []byte) ([]byte, error) {
	block, _ := pem.Decode(r.pubKey)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

func (r *decryptor) DecryptString(ciphertext string) (string, error) {
	cipherBytes, err := base64.RawURLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	plainText, err := r.Decrypt(cipherBytes)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func (r *decryptor) Decrypt(cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(r.priKey)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
