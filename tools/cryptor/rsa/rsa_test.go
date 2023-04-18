package rsa

import (
	"encoding/hex"
	"testing"
)

const pubKey = `-----BEGIN NOX RSA PUBLIC KEY -----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4SYDI3Rot/FNuZQgSVJr
Z6HrDuk9+yNIXNBWAv/tH4NzvGchI+Kzh2ZsVjmBRpFDsOxgMlmwhx70xdh4yWem
Sf8pYzQ6YTmUaQtcM+pyFaJ9k3SAVITEJrosWmrAUKTGiEwvUO1whhnwiNpiz18b
uDHqgO9j7fujsSjqtbuG+XBiB2jpx+7h3X7lDQJL+UzhY0DZ6vE5VRoDZO/ihZiz
AqHdKyX5e+qf2InysdEczlhsFGND8UPcc7S4lBnswbxsX11BJV5dOGgHuawPOeSg
Vt5xPXShgINdHNyiR1iBrOnIeQqwRrdl0lwjtdEo6E8Q9r+Q3HvYS1ddDOuHWjOM
6QIDAQAB
-----END NOX RSA PUBLIC KEY -----
`

const priKey = `-----BEGIN NOX RSA PRIVATE KEY -----
MIIEpAIBAAKCAQEA4SYDI3Rot/FNuZQgSVJrZ6HrDuk9+yNIXNBWAv/tH4NzvGch
I+Kzh2ZsVjmBRpFDsOxgMlmwhx70xdh4yWemSf8pYzQ6YTmUaQtcM+pyFaJ9k3SA
VITEJrosWmrAUKTGiEwvUO1whhnwiNpiz18buDHqgO9j7fujsSjqtbuG+XBiB2jp
x+7h3X7lDQJL+UzhY0DZ6vE5VRoDZO/ihZizAqHdKyX5e+qf2InysdEczlhsFGND
8UPcc7S4lBnswbxsX11BJV5dOGgHuawPOeSgVt5xPXShgINdHNyiR1iBrOnIeQqw
Rrdl0lwjtdEo6E8Q9r+Q3HvYS1ddDOuHWjOM6QIDAQABAoIBAE/2nC6/bWHHBf/f
KAhiiVrpGv7Uv/qt8UllhObT1nfWzAgh6JdCMNjp+5g0HEHf3da8XP4E9LlIuU57
L4r4kQq+4QgmyIU3Wgpkyn51ycb9SvRP8DQUY0YN4SrLtzntC8XkqPlpGwnFtLQm
O+nCsamdsNfSLf+b2+tTyOh6jdmQ8ebQDKFtMDMXmEXviNaEeohAkBiZyHBg02eL
kyINyMQ9fm1ehhf+qpspuaLqLoNI5bJMFRT3bDCm2tssUCDtshnZ0i0uszneVUn5
6b/sqTgn2bDFsz5/feFQTXXtc+aLZHu2z75NyFUCbN2pm2X/7TUH2IXc8w88ulAC
sV8kIcECgYEA6ZqXPYOb75uRHgu3j6AIyBF6IEfZwdXyvF9WCOZZRRgl05vqBsfl
La1uhgTP2Tu4RrlbcOEBYI1Xr4LszNVqC5uQ+Q0uxZuz5kdbc09MDKKIHGgbAn1y
L0fHicNnhQrFjS/qlfjuFoxFvFL+bmWVUJdbAJcMsSDdz66ndq4oQJUCgYEA9rvm
ARIbaIKDlVG2+Vh1qxtdqgGoXvfnDIGK82JwL/6QmNwf+VGOBKXNXuYRIUkTDAYl
nXevAQY5trt8RvnaEMoAgIFo0t2TzfFgtrt/AQUkYMr5aA/AUHJWFo4dqhnbQayw
77SgitrM3pu5VRv2pcoyvteFfeAr7ho9ifFcogUCgYEAlmqjTcmro8pA097pkEKU
xPZz88swDN7NULv2cv6XpqTY3nu7Yiheil3tF8CLcS5CBtAdb/6B24DHYEzmWzJj
+Rtvat4xKq6KVRHHceUya8RP7jKeiajq2ETY1/0JcCIyeCnNXEoQTFaAJLYv/DfO
wDjlMOkSFA+8o5irXgDgDE0CgYEAz2NUiuXbX9cEHwhWUfPGxBH2gqVuDpU1AdAJ
KVY+kjjY9YNg5MFmLAXpP7EGCQJRHjQ0fc1Rrfy7mhqw9T+hyabhESK1JzjTubD/
59hlgkf+MdWqGYvTzEqujbrauFWPc/5horoij1QmC569AwrEIMp6KOsi65D3m+U2
gLH9eLkCgYB2bCi1OxDO8nFBGwobpESPQrT5w3yXKwv8p2pHrJ8xUdVaYGZEEn3j
zGvOJ9j9Io4a/lPDa14dotnfeEU/so/FKtj12t02xY39emqWcQBUTkS2lH2Y3wTr
OLRCrtrMORFhOvfGwiorbh//BNTuzQjotjaveqOZdOtSEbCl+PRAHA==
-----END NOX RSA PRIVATE KEY -----
`

func Test_RsaCrypt(t *testing.T) {
	res, err := NewEncryptor(pubKey).Encrypt([]byte("hello gopher"))
	if err != nil {
		t.Errorf("rsa.Encrypt() error = %v", err)
		return
	}
	result := hex.EncodeToString(res)
	t.Logf("rsa.Encrypt() = %v", result)

	val, _ := hex.DecodeString(result)
	res, err = NewDecryptor(priKey).Decrypt(val)
	if err != nil {
		t.Errorf("rsa.Decrypt() error = %v", err)
		return
	}
	t.Logf("rsa.Decrypt() = %v", string(res))
}

func Test_RsaCryptString(t *testing.T) {
	res, err := NewEncryptor(pubKey).EncryptString("hello gopher")
	if err != nil {
		t.Errorf("rsa.Encrypt() error = %v", err)
		return
	}
	t.Logf("rsa.Encrypt() = %v", res)

	res, err = NewDecryptor(priKey).DecryptString(res)
	if err != nil {
		t.Errorf("rsa.Decrypt() error = %v", err)
		return
	}
	t.Logf("rsa.Decrypt() = %v", string(res))
}
