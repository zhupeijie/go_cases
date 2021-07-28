package resh

import (
	"io"
	"encoding/base64"
	"io/ioutil"
	"bytes"
)


func DecryptByPublicKey(encryptStr string, pubKey string) (decryStr string, err error) {
	var in []byte
	in, err = base64.StdEncoding.DecodeString(encryptStr)

	if err != nil {
		return "111", err
	}

	out := bytes.NewBuffer(nil)
	inBytes := bytes.NewReader(in)
	//err := this.IO(bytes.NewReader(in), out, mode)
	//if err != nil {
	//	return nil, err
	//}
	//return ioutil.ReadAll(out)
	publicKey,_ := getPubKey([]byte(pubKey))

	k := (publicKey.N.BitLen() + 7) / 8
	println(k)
	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = inBytes.Read(buf)
		if err != nil {
			if err == io.EOF {
				return "222", err
			}
			return "333", err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}

		b, err = pubKeyDecrypt(publicKey, b)
		if err != nil {
			return "22", err
		}
		if _, err = out.Write(b); err != nil {
			return "333", err
		}
	}
	resB,err := ioutil.ReadAll(out)
	return string(resB), err

}

/*
func encryptByPrivateKey(sourceStr string,privkey string)(encryStr string,err error)  {
	
}
*/
