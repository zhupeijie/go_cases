package test

const CV0  = `

`


//
//func main() {
//
//	arr := make(map[string]string)
//
//	arr["cv0"] = ``
//	arr["cv1"] = ``
//	arr["cv2"] = ``
//	arr["cv3"] = ``
//
//
//	str := "Kge87hIpNBW9Dmw7RDYMcizulYYI0iu7kDgwVulC5p1-JrFdOv5OkoMNRb0V52-I1gcj0LmTQxLWsdrsae5wptXqespmmraZYDd3ndztG4zuXWnaJD2pi0oKm5SjdEcKu66AEPjMM7J0xDDUYX648CZPKTtlzYvMq7xdmSNlESU.cv0"
//	strArr := strings.Split(str, ".")
//
//
//	fmt.Println(len(strArr))
//
//	strA := strArr[0]
//	strLen := 4 - len(strA)%4
//	strArr[1] = strings.Repeat("=", strLen)
//
//	strB := strings.Join(strArr, "")
//	fmt.Println(strB)
//	repNu := strings.Count(strB, "-_")
//	strB = strings.Replace(strB, "-_", "+/", repNu)
//
//	fmt.Println(string(strB[:]))
//
//
//	repNu = strings.Count(strB, "-")
//	strB = strings.Replace(strB, "-", "+", repNu)
//	fmt.Println(string(strB[:]))
//
//
//	strC, _ := base64.StdEncoding.DecodeString(strB)
//	certV0 := `
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDdB9VZL49+HD8Oo6/yh63tBe9D
//sNkyAGaNhnJmJS+irCYbjuFnzzu9zmfkA8eUwp9m9Uez8B8706gLgXQLlehkf8VJ
//Oitp52Zt5PwjLevPlXvE6VfRCMdvJ9wL5P7Lb5UXW2J9LztJ8sE7eYCMpmgHPXas
//BToQRInmKfXVJ9d46wIDAQAB
//-----END PUBLIC KEY-----
//`
//	buf, e := rsa.PublicDecrypt(strC, certV0, rsa.RSA_PKCS1_PADDING)
//	if e == nil {
//		fmt.Printf("Decrypt: %s", string(buf))
//	} else {
//		fmt.Printf("%s\n", e.Error())
//		return
//	}
//
//	//decStr, decErr := encrypt.RsaDecrypt(strC, certV0)
//
//
//}
