package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 5}
	val := 4
	res := delSliSame(nums, val)
	for i := 0; i < res; i++ {
		fmt.Println(nums[i])
	}

}

func delSliSame(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func delSliSameV1(nums []int, val int) int {
	i := len(nums)
	j := 0
	for {
		if j >= i {
			return j
		}
		if nums[j] == val {
			nums[j] = nums[i-1]
			i--
		} else {
			j++
		}
	}
}
