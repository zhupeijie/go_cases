package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := "A man, a plan, a canal: Panama"
	res :=  isPalindrome(s)
	fmt.Println(res)
}
//import "unicode"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	fmt.Println(s,len(s),isAlNum(s[6]))
	r := len(s)-1
	l := 0
	for l<r {
		if  l<r && !isAlNum(s[r]){
			r--
			continue
		}
		if l<r && !isAlNum(s[l]) {
			l++
			continue
		}
		if l < r {
			if s[l] != s[r]{
				fmt.Println(l,r, "===",string(s[l]),"---",string(s[r]))
				return false
			}
			l++
			r--
		}
	}
	return true

}

func isAlNum(ch byte)bool{
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
