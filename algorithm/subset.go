package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	resp := make([][]int, 0)
	mask := 1<<len(nums) - 1     // 标识位
	for i := 0; i <= mask; i++ { // 循环到mask,mask的二进制位为1,代表要选择指定位置的值
		index := 0
		tmp := []int{}
		for j := i; j > 0; j = j >> 1 {
			if j&1 == 1 {
				tmp = append(tmp, nums[index])
			}
			index++
		}
		resp = append(resp, tmp)
	}
	return resp
}
