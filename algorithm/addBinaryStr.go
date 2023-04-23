package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "11"
	s2 := "10011"
	fmt.Println(addBinary(s1, s2))
}

func addBinary(a string, b string) string {

	var tmp string
	if len(a) < len(b) {
		tmp = a
		a = b
		b = tmp
		tmp = ""
	}
	t := 0
	for i := 0; i < len(a); i++ {
		if i < len(a) {
			t += int(a[len(a)-i-1] - '0')
		}
		if i < len(b) {
			t += int(b[len(b)-i-1] - '0')
		}
		tmp = strconv.Itoa(t%2) + tmp
		t = t/2
	}

	if t == 1 {
		tmp = "1" + tmp
	}
	return tmp
}
