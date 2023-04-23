package aes

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_CbcCrypt(t *testing.T) {
	c := NewCbc("hellogopherhello")
	res, err := c.Encrypt([]byte("hello gopher"))
	if err != nil {
		t.Errorf("aescbc.Encrypt() error = %v", err)
		return
	}
	result := hex.EncodeToString(res)
	t.Logf("aescbc.Encrypt() = %v", result)

	val, _ := hex.DecodeString(result)
	res, err = c.Decrypt(val)
	if err != nil {
		t.Errorf("aescbc.Decrypt() error = %v", err)
		return
	}
	t.Logf("aescbc.Decrypt() = %v", string(res))
}

func Test_CbcCryptString(t *testing.T) {
	c := NewCbc("hellogopherhello")
	res, err := c.EncryptString("hello gopher")
	if err != nil {
		t.Errorf("aescbc.Encrypt() error = %v", err)
		return
	}
	t.Logf("aescbc.Encrypt() = %v", res)

	res, err = c.DecryptString(res)
	if err != nil {
		t.Errorf("aescbc.Decrypt() error = %v", err)
		return
	}
	t.Logf("aescbc.Decrypt() = %v", res)
}

func Test_DecryptPhone(t *testing.T) {
	c := NewCbc("538a8ae285fcf3ee61e910be9dc2fe4d")
	enPhoneSlice := []string{
		"ah1V6q4HBRzQoJh4hIiUGg==",
		"3YIELIkq0sL4jm5GwiqVoA==",
		"tpN3ki/dginv5RTXbVJv3w==",
		"GdC9k5+KI8ykZFknBkNnOw==",
		"WycKFtdSBDGTY7RzFBkWCQ==",
		"dR4uX33b6rowGiY1lBfb1Q==",
	}
	for _, v := range enPhoneSlice {
		phone, _ := c.DecryptString(v)
		fmt.Printf("\"%s\" %s\n", v, phone)
	}
}
