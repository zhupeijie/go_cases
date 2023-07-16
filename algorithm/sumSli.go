package main

import (
	"fmt"
	"sort"
)

/*
输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
*/

func main() {
	in := []int{2, 5, 2, 1, 2}
	target := 5
	// rsp := sumEqual(in, target)
	sort.Ints(in)
	rsp := sumEqualAll(in, target)
	mdm := make(map[string]bool)
	for _, v := range rsp {

	}

	fmt.Println(rsp)
}

func sumEqual(in []int, target int) [][]int {
	resArr := make([][]int, 0)
	sort.Ints(in)
	var tg int
	for i := 0; i < len(in); i++ {
		if in[i] > target {
			break
		}
		tg = target
		tmp := make([]int, 0)
		for j := i; j < len(in); j++ {
			tg -= in[j]
			tmp = append(tmp, in[j])
			if tg == 0 {
				resArr = append(resArr, tmp)
			}
			if in[j] > tg {
				break
			}
		}
	}
	return resArr
}

//  穷举
func sumEqualAll(in []int, target int) [][]int {
	resArr := make([][]int, 0)
	for i := 0; i < len(in); i++ {
		if in[i] > target {
			continue
		}
		if in[i] == target {
			resArr = append(resArr, []int{target})
			return resArr
		}
		if target == 0 {
			return resArr
		}
		tmpArr := sumEqualAll(in[1:], target-in[0])
		for _, v := range tmpArr {
			v = append(v, in[i])
			resArr = append(resArr, v)
		}
	}
	return resArr
}
