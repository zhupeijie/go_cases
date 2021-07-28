package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	//"mississippi"
	//"issip"
	//hay := "asfaeqieisis"
	//ne := "aeq"
	hay := "amississippi"
	ne := "issip"
	//res := searchNeedle(hay, ne)
	res := strStr1(hay, ne)
	//res := strstr2(hay, ne)
	fmt.Println(res)


}

func strstr2(haystack string, needle string)int  {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0  {
		return 0
	}
	if l1 == 0 || l1<l2 {
		return -1
	}
	for i:=0;i<l1-l2;i++{
		if  haystack[i:i+l2]==needle{
			return i
		}
	}
	return -1
}

//暴力破解
func searchNeedle(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			same := 0
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					same = 1
					break
				}
			}
			if same == 0 {
				return i
			}
		}
		continue
	}
	return -1
}

func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	aN := 26
	//module := int(math.Pow(2,31))
	module := int(math.Pow(2,31))
	//f(n+1) = fn
	h := 0
	needle_h := 0
	al :=1
	for i := 0;i<len(needle);i++ {
		h = (h*aN+calVal(i,haystack))%module
		needle_h = (needle_h*aN+calVal(i,needle))%module
		al = (al*aN)%module
	}
	if needle_h == h {
		return 0
	}

	// h+1  = (h-(c-1)*al)*an+now*al
	for i:=1;i<len(haystack)-len(needle)+1;i++ {
		h = (h*aN+calVal(i+len(needle)-1,haystack)-calVal(i-1,haystack)*al)%module
		if h == needle_h {
			return i
		}
	}

	return -1
}

func calVal(index int, s string) int {
	a := int(s[index]) - int("a"[0])
	return a
}

func searchNeedleInt(sint []int,needle int)int  {
	for k,v:= range sint{
		if v==needle{
			return k
		}
		if needle < sint[0]{
			return 0
		}
		if needle > sint[len(sint)-1] {
			return len(sint)
		}
		if v<needle && sint[k+1]>needle {
			return k+1
		}
	}
	return 0
}

type JsA struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
	Data  JsData `json:"data"`
}

type JsData struct {
	KnowledgeList []JsKInfo `json:"knowledge_list"`
}

type JsKInfo struct {
	KnowledgeId int `json:"knowledge_id"`
	KnowledgeName string `json:"knowledge_name"`
}




func jsUnm ()  {
	a := `{"errcode":0,"errmsg":"\u8bf7\u6c42\u6210\u529f","data":{"knowledge_list":[{"knowledge_id":3,"knowledge_name":"\u77e5\u8bc6\u70b93","knowledge_study_level":"58%","knowledge_test_rate":"50%"},{"knowledge_id":2,"knowledge_name":"\u77e5\u8bc6\u70b92","knowledge_study_level":"61%","knowledge_test_rate":"40%"},{"knowledge_id":4,"knowledge_name":"\u77e5\u8bc6\u70b94","knowledge_study_level":"60%","knowledge_test_rate":"30%"},{"knowledge_id":1,"knowledge_name":"\u77e5\u8bc6\u70b91","knowledge_study_level":"85%","knowledge_test_rate":"30%"},{"knowledge_id":5,"knowledge_name":"\u77e5\u8bc6\u70b95","knowledge_study_level":"85%","knowledge_test_rate":"20%"}]}}`
	b := new(JsA)

	err := json.Unmarshal([]byte(a),&b)
	fmt.Println(err)
	fmt.Println(b)

}