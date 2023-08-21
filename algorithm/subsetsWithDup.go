package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{4, 4, 4, 1, 4}
	nums := []int{1, 2, 2}
	resp := subsetsWithDup(nums)
	fmt.Println(resp)
	//nums1 := [][]int{{}, {4}, {4, 4}, {4, 4, 4}, {1}, {4, 1}, {4, 4, 1}, {4, 4, 4, 1}}
	//a := newD(nums1, 4)
	//fmt.Println(nums1)
	//fmt.Println(a)
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	resp := make([][]int, 0)
	mask := 1<<len(nums) - 1
	for i := 0; i <= mask; i++ { // 循环子集的长度
		index := 0
		tmp := []int{}
		needAppend := true
		for j := i; j > 0; j = j >> 1 {
			if j&1 == 1 {
				if index > 0 && nums[index] == nums[index-1] && i>>(index-1)&1 == 0 { // 上个等值元素未被选中,那么跳过
					index++
					needAppend = false
					break
				}
				tmp = append(tmp, nums[index])
			}
			index++
		}
		if needAppend {
			resp = append(resp, tmp)
		}
	}

	return resp
}

func newD(in [][]int, num int) [][]int {
	resp := [][]int{}
	for _, v := range in {
		//tmp := append(v[:], num)
		tmp := append([]int{num}, v...)
		exist := false
		for i := len(in) - 1; i >= 0; i-- {
			if compare(in[i], tmp) {
				exist = true
				break
			}
		}
		if !exist {
			for i := len(resp) - 1; i >= 0; i-- {
				if compare(resp[i], tmp) {
					exist = true
					break
				}
			}
		}

		if !exist {
			resp = append(resp, tmp)
		}
	}
	return resp
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	tmp := map[int]int{}
	for i, v := range a {
		tmp[v]++
		tmp[b[i]]--
	}
	for _, v := range tmp {
		if v != 0 {
			return false
		}
	}
	return true
}
