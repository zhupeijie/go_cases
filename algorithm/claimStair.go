package main

import "fmt"

func main() {
	//
	fmt.Println(claim(44))}

func claim(n int) int {
	l,r,t := 0,0,0
	for i := 0; i < n; i++ {
		if i == 0 {
			l = 1
			continue
		}
		r = l+t
		t = l
		l = r

	}
	return l+t
	////1,找出有多少中情况
	////2. 计算没种情况的排列组合数量  f(0) = 1,f(1) =1
	//if n == 0 || n==1{
	//	return 1
	//}
	//
	//return claim(n-1) + claim(n-2)
}

