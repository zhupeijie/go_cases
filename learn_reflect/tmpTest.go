package main

import (
	"fmt"
	"math"
)

type T1 struct {
	Name string
	Age  int
}

func main() {
	//x := 123
	//v := reverse(x)
	//fmt.Println(v)
	x := Hw(1001)
	fmt.Println(x)
}

func Hw(x int) bool {

	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	rev := 0
	if x < 100 {
		if x%10 == x/10 {
			return true
		}
		return false
	}
	for {
		pop := x % 10
		rev = rev*10 + pop
		x /= 10
		fmt.Println(rev)
		fmt.Println(x)
		if x == rev || x/10 == rev {
			return true
		}
		if x < rev {
			return false
		}
	}
}

func test1() {

}

func test2(v T1) (error, ) {
	v = T1{}
	return nil
}
func reverse(x int) int {
	var rev int
	max := int(math.Pow(2, 31) - 1)
	min := int(math.Pow(-2, 31))
	fmt.Println(max)
	fmt.Println(min)

	for {
		pop := x % 10
		x /= 10

		if rev > max/10 || (x > 7 && x == max/10) {
			rev = 0
			break
		}
		if rev < min/10 || (x < -8 && x == (min/10)) {
			rev = 0
			break
		}
		rev = rev*10 + pop
		if x == 0 {
			return rev
		}
	}
	return rev

}
