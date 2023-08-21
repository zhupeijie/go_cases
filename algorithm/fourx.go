package main

import (
	"fmt"
)

func main() {
	fmt.Println(fourx2(16))
}

func fourx(n int) bool {
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func fourx2(n int) bool {
	tmp := make(map[int32]bool)
	for i := 0; i < 15; i++ {

		tmp[1<<(i*2)] = true
	}
	fmt.Println(tmp)
	return false
}
