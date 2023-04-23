package main

import (
	"fmt"
	"sort"
)
/**
给定一个数 n，如 23121；给定一组数字 A 如 {2,4,9}，求由 A 中元素组成的、小于 n 的最大数，如小于 23121 的最大数为 22999
 */
func main() {
	target := 11121
	in := []int{2, 4, 9}
	res := leesThan(in, target)
	fmt.Println(res)
}

// 拆分数字为数组
func getEvery(in int) []int {
	res := make([]int, 0)
	for in > 0 {
		if in < 10 && in > 0 {
			res = append(res, in)
			in = 0
			return res
		}
		res = append(res, in%10)
		in = in / 10
	}
	return res
}

// 给一个目标数,和一个数组,找出小于目标数最大的数
func leesThan(in []int, target int) int {
	sort.Ints(in) // 先排序
	lessthen := 0
	res := 0
	targetSli := getEvery(target)
	for i := len(targetSli) - 1; i >= 0; i-- {
		fmt.Println("----", res)
		if lessthen == 1 {
			res = res*10 + in[len(in)-1]
			continue
		}
		j := sort.SearchInts(in, targetSli[i])
		if in[j] != targetSli[i] {
			lessthen = 1
			if j == 0 && res == 0 {
				continue
			}
			res = res*10 + in[j-1]
			continue
		}
		res = res*10 + in[j]
	}
	return res
}
