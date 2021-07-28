package main

import "fmt"

func main() {
	res := fun1("([)]")
	fmt.Println(res)
}

func fun1(s string) bool {
	if len(s) <= 1 {
		return false
	}
	leftS := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if left(s[i]) {
			leftS = append(leftS, s[i])
		}
		if right(s[i]) {
			if len(leftS) == 0 || getRight(leftS[len(leftS)-1]) != s[i] {
				return false
			}
			leftS = leftS[:len(leftS)-1]
		}
		fmt.Println(leftS, left(s[i]), s[i])
	}
	return true
}

func test(s string) {
	fmt.Println(s)
	fmt.Println(s[0 : len(s)-1])
}

//
func left(s uint8) bool {
	switch s {
	case 40,123,91:
		return true
	default:
		return false
	}
}
func right(s uint8) bool {
	switch s {
	case 41,125,93:
		return true
	default:
		return false
	}
}

func getRight(v uint8) uint8 {
	switch v {
	case 40:
		return 41
	case 123:
		return 125
	case 91:
		return 93
	}
	return 0
}
