package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1}}, 1))
}
func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	i := 0
	// 先找道所在的行
	for ; i < len(matrix); i++ {
		if target < matrix[i][0] {
			i = i - 1
			break // 递增的值,所以,第一个小于的既是
		}
	}
	if i == len(matrix) {
		i = len(matrix) - 1
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[i][j] == target {
			return true
		}
	}
	return false
}
