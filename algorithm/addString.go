package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	licenseKeyFormatting("5F3Z-2e-9-wf", 2)
	//adds := addStrings("123", "123")
	//fmt.Println(adds)
}
func addStrings(num1, num2 string) string {
	ans := ""
	pre := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {

		if i >= 0 {

			pre = pre + int(num1[i]-'0')
		}
		if j >= 0 {
			pre = pre + int(num2[j]-'0')
		}

		ans = strconv.Itoa(pre%10) + ans
		pre = pre / 10
	}
	if pre > 0 {
		ans = strconv.Itoa(pre) + ans
	}
	return ans
}

func licenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(s, "-", "")

	preNum := len(s) % k
	if preNum == 0 {
		preNum = k
	}
	s1 := ""
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			s1 += string(unicode.ToUpper(rune(s[i])))
			continue
		}
		s1 += string(s[i])
	}
	fmt.Println(s, s1)
	s = s1
	s1 = ""
	fmt.Println(s, s1)
	s1 += s[:preNum]
	start := preNum
	for ; start <= len(s)-k; start = start + k {
		s1 += "-" + s[start:start+k]
	}
	fmt.Println(s, s1)
	return s1

}
