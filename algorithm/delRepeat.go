package main

import "fmt"

func main() {
	//in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//in := []int{1, 1, 2, 2}
	in := []int{1, 1, 1,1,1}
	res := delSame(in)
	for l := 0; l < res; l++ {
		fmt.Println(in[l])
	}
}

func delSame(nums []int) int {
	num := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		if i == num-1 {
			break
		}
		if nums[i] == nums[i+1] {
			for j := i+1;j<len(nums)-1;j++ {
				if nums[i] == nums[j]  {
					num--
					tmp :=nums[i+1]
					for l := i+1;l< len(nums)-1;l++ {
						nums[l] = nums[l+1]
					}
					nums[len(nums)-1] = tmp
				}
			}
		}
		if i == num-1 {
			break
		}
	}
	return num
}
