package aes

import (
	"encoding/hex"
	"testing"
)

func Test_CtrCrypt(t *testing.T) {
	c := NewCtr("hellogopherhello")
	res, err := c.Encrypt([]byte("hello gopher"))
	if err != nil {
		t.Errorf("aesctr.Encrypt() error = %v", err)
		return
	}
	result := hex.EncodeToString(res)
	t.Logf("aesctr.Encrypt() = %v", result)

	val, _ := hex.DecodeString(result)
	res, err = c.Decrypt(val)
	if err != nil {
		t.Errorf("aesctr.Decrypt() error = %v", err)
		return
	}
	t.Logf("aesctr.Decrypt() = %v", string(res))
}

func Test_CtrCryptString(t *testing.T) {
	c := NewCtr("hellogopherhello")
	res, err := c.EncryptString("hello gopher")
	if err != nil {
		t.Errorf("aesctr.Encrypt() error = %v", err)
		return
	}
	t.Logf("aesctr.Encrypt() = %v", (res))

	res, err = c.DecryptString(res)
	if err != nil {
		t.Errorf("aesctr.Decrypt() error = %v", err)
		return
	}
	t.Logf("aesctr.Decrypt() = %v", res)
}
