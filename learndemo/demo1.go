package main

import (
	"fmt"
)

func main() {
	str := "12345678"
	str = Commav2(str)
	fmt.Println(str)
}

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func Commav2(in string) string {
	byteA := []byte(in)
	n := len(byteA)
	if n <= 3 {
		return in
	}
	// bytes.Index()
	newBytes := []byte{}
	count := len(byteA)
	for i := 0; i < count; i++ {
		newBytes = append(newBytes, byteA[i])
		if (count-i)%3 == 1 && i != count-1 {
			newBytes = append(newBytes, ',')
		}
	}
	return string(newBytes)
}
