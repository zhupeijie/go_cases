package main

import "fmt"

func main() {
	// [[2],[3,4],[6,5,7],[4,1,8,3]]
	data := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	resp := minimumTotal(data)
	fmt.Println(resp)
}

func minimumTotal(triangle [][]int) int {

	//minCount := make([]int, len(triangle[len(triangle)-1]))
	for i, nums := range triangle {
		if i == 0 {
			continue
		}
		for j, num := range nums {
			if j == 0 {
				triangle[i][0] = triangle[i-1][0] + num
			} else if j == i {
				triangle[i][j] = triangle[i-1][j-1] + num
			} else {
				triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + num
			}
		}
		//fmt.Println(minCount)
	}
	minN := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		minN = min(minN, v)
	}
	return minN
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
