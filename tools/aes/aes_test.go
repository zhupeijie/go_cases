package aes

import (
	"fmt"
	"testing"
)

var key = "ndfuewnfkansdvniuebvkasdbviegr1a"

func TestAesEncrypt(t *testing.T) {
	str := "hello aes "
	ens, err := Encrypt([]byte(str), key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ens, err)
}

func TestAesDecrypt(t *testing.T) {
	str := "bu3PAEJN2hFq05tNRlA6PA=="

	des, err := Decrypt(str, key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(des), err)
}
