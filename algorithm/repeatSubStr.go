package main

import "fmt"

func main() {
	fmt.Println(repeatedSubStr("abac"))
}
func repeatedSubStr(s string) bool {
	if len(s) <= 1 {
		return false
	}
	n := len(s)
	for i := 1; i < len(s)/2+1; i++ { // 子串的长度
		if n%i > 0 { // 不能整除的都不对
			continue
		}
		ok := true
		for j := 1; j <= n/i; j++ {
			fmt.Println(s[(j-1)*i:j*i], s[:i])
			if s[(j-1)*i:j*i] != s[:i] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
