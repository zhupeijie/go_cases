package main

import "fmt"

func main() {
	s := "badc"
	b := "baba"
	fmt.Println(isIsomorphic(s, b))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}
	tmpM := make(map[byte]byte)
	tmpM1 := make(map[byte]bool)
	for i := range s {
		if rv, ok := tmpM[s[i]]; ok {
			if rv != t[i] {
				return false
			}
			continue
		}
		if _, ok := tmpM1[t[i]]; ok {
			return false
		}
		tmpM[s[i]] = t[i]
		tmpM1[t[i]] = true
	}
	return true
}
