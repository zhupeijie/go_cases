package test

import (
	"encoding/base64"
	"crypto/x509"
	"crypto/sha1"
	"crypto/rsa"
	"crypto"
)


/**
xes_rfh=gHVM1hKJx11-a0Gn8LBVPk9_W-xgxuX-oI73DXtgHVotHx2pVrKC0asq05TPp_RescEVKmlArkbRLthUAoKcbHE3C297Kax8-Bzmx9Qwd3Nsr4pv3lIMUf-RV_ETDwUuXkEKJsDPizWNwyIvv9JwqdER1B8H7StnLRTVCNXJwVs.cv1


gHVM1hKJx11-a0Gn8LBVPk9_W-xgxuX-oI73DXtgHVotHx2pVrKC0asq05TPp_RescEVKmlArkbRLthUAoKcbHE3C297Kax8-Bzmx9Qwd3Nsr4pv3lIMUf-RV_ETDwUuXkEKJsDPizWNwyIvv9JwqdER1B8H7StnLRTVCNXJwVs.cv1

 */

func RsaVerySignWithSha1Base64(originalData, signData, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}
//
//func main() {
//	originalData := ""
//	hash := sha1.New()
//	hash.Write([]byte(originalData))
//	println(originalData)
//	println("exit")
//	/*
//	signData := "gHVM1hKJx11-a0Gn8LBVPk9_W-xgxuX-oI73DXtgHVotHx2pVrKC0asq05TPp_RescEVKmlArkbRLthUAoKcbHE3C297Kax8-Bzmx9Qwd3Nsr4pv3lIMUf-RV_ETDwUuXkEKJsDPizWNwyIvv9JwqdER1B8H7StnLRTVCNXJwVs="
//
//	sign, err := base64.StdEncoding.DecodeString(signData)
//	if err != nil {
//		println(err.Error())
//		print("this is error 1")
//		return
//	}
//	pubKey := `
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDdB9VZL49+HD8Oo6/yh63tBe9D
//sNkyAGaNhnJmJS+irCYbjuFnzzu9zmfkA8eUwp9m9Uez8B8706gLgXQLlehkf8VJ
//Oitp52Zt5PwjLevPlXvE6VfRCMdvJ9wL5P7Lb5UXW2J9LztJ8sE7eYCMpmgHPXas
//BToQRInmKfXVJ9d46wIDAQAB
//-----END PUBLIC KEY-----
//`
//	public, _ := base64.StdEncoding.DecodeString(pubKey)
//	pub, err1 := x509.ParsePKIXPublicKey(public)
//	if err1 != nil {
//		print("this is error 2")
//
//		return
//	}
//	originalData := ""
//	hash := sha1.New()
//	hash.Write([]byte(originalData))
//	print(originalData)
//	rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
//	return
//	*/
//}
