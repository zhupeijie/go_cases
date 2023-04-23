package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

func Sign1() {
	mircTime := 1663207990810
	secret := "********"
	str := fmt.Sprintf("%d\n%s", mircTime, secret)
	fmt.Println(str)

	// s256 := sha256.New()
	// _,_ = io.WriteString(s256, str)

	strSha := hmac.New(sha256.New, []byte(secret))
	_, err := strSha.Write([]byte(str))
	if err != nil {
		fmt.Println("some wrong has happen ,", err)
		return
	}
	sig1 := base64.StdEncoding.EncodeToString(strSha.Sum(nil))
	fmt.Println("1111-----", sig1)
	need := url.QueryEscape(sig1)
	// 此方法等同于 PHP代码如下 base64_encode(hash_hmac('sha256', $data, $app_secret, false));
	fmt.Println(" need str sing ====", need)
}

func Sign11() {

	secret := "this is secret"
	str := fmt.Sprintf("%d\n%s", time.Now().UnixNano()/1e6, secret)
	strSha := hmac.New(sha256.New, []byte(secret))
	strSha.Write([]byte(str))
	// 这里等同于PHP代码如下 base64_encode(hash_hmac('sha256', $data, $app_secret, true));

	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(strSha.Sum(nil)))
	fmt.Println(sign)
}
