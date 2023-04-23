package main

import (
	"fmt"
	"math"
)

func main (){
	fmt.Println('.')
	s := myAtoi("   -42")
	fmt.Println(s)
}

func myAtoi(s string) int {
	resNum := 0
	tmp := make([]int32, 0)
	start := true
	for _, i := range s {

		if start && i == ' ' {
			continue
		}
		if start && (i == '-' || i == '+') {
			tmp = append(tmp, i)
			start = false
			continue
		}
		fmt.Println("---",i,start)
		if start && i > '9' || i < '0' && i != '-' && 1 != '+' && i != ' ' {
			break
		}
		fmt.Println("@@@@",i,start)
		if !start && i > '9' || i < '0' {
			break
		}
		if i >= '0' && i <= '9' {
			tmp = append(tmp, i)
			start = false
		}
	}
	min := math.MinInt32
	max := math.MaxInt32
	f := 1
	for _, b := range tmp {
		if b == '-' {
			f = -1
		}
		if b == '-' || b == '+' {
			continue
		}
		tmp := 10*resNum + int(b) - 48
		if f*tmp < min {
			return min
		}
		if f*tmp > max {
			return max
		}
		resNum = tmp
	}
	return f * resNum
}
