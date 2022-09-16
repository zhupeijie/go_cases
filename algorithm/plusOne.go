package main

import "fmt"

func main()  {
	in := []int{9}
	fmt.Println(plusOne(in))
}
func plusOne(digits []int) []int {
	tmp := 1
	for i := len(digits)-1; i >=0 ; i-- {
		if digits[i]+ tmp == 10{
			digits[i] = 0
		}else{
			digits[i] +=tmp
			return digits
		}
	}
	res := make([]int,len(digits)+1)
	res[0] = tmp
	return res
}
