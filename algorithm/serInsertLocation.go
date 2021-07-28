package main

import "fmt"

/**
[1,3,5,6]
5
[1,3,5,6]
2
[1,3,5,6]
7
[1,3,5,6]
0
[1]
0
=======
2
1
4
0
0
*/
func main() {
	//res := serInsertTwo([]int{1,3,5,6},5)
	//res := serInsertTwo([]int{1,3,5},5)
	//res := serInsert([]int{1}, 1)
	//fmt.Println(res)
}

func serInsert(nums []int, target int) int {
	for k, v := range nums {
		if target <= v {
			return k
		}
	}
	return len(nums)
}

//二分法查找
func serInsertTwo(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	return serInsertMid(nums, target, 0, len(nums)-1)
}

func serInsertMid(nums []int, target int, start int, end int) int {
	middle := (start + end) / 2
	fmt.Println(nums,start,end,middle)
	if middle == 0 {
		return 1
	}
	if nums[middle] >= target && nums[middle-1] < target {
		fmt.Println(nums,start,end,middle,"=====")
		return middle
	}
	if nums[middle] > target {
		return serInsertMid(nums, target, start, middle)
	}
	if nums[middle] < target && nums[middle+1] >= target {
		return middle+1
	}
	if nums[middle] < target {
		return serInsertMid(nums, target, middle, end)
	}
	fmt.Println(nums,start,end,middle,"++++")
	return end
}
