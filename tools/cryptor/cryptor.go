package cryptor

type Cryptor interface {
	// EncryptString encrypt
	EncryptString(plainText string) (string, error)

	// Encrypt encrypt
	Encrypt(plainText []byte) ([]byte, error)

	// DecryptString decrypt
	DecryptString(cipherText string) (string, error)

	// Decrypt decrypt
	Decrypt(cipherText []byte) ([]byte, error)
}
