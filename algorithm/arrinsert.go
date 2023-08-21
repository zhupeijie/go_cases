package main

import "fmt"

func main() {
	a := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	fmt.Println(a)
	//[[1,2],[3,5],[6,7],[8,10],[12,16]]
	//	[4,8]
}
func insert(intervals [][]int, newInterval []int) [][]int {
	resp := make([][]int, 0)
	if newInterval[1] < intervals[0][0] {
		resp = append(resp, newInterval)
		resp = append(resp, intervals...)
		return resp
	}
	for _, v := range intervals {
		if v[1] < newInterval[0] || v[0] > newInterval[1] {
			resp = append(resp, v)
			continue
		}
		start := min(v[0], newInterval[0])
		end := max(v[1], newInterval[1])
		newInterval = []int{start, end}
		if len(resp) > 0 && resp[len(resp)-1][1] >= start {
			resp[len(resp)-1][1] = end
			continue
		}
		resp = append(resp, newInterval)
	}
	return resp
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateMatrix(n int) [][]int {

	resp := make([][]int, n)
	for i := range resp {
		resp[i] = make([]int, n)
	}
	// 四个方向, 右,下,左,上 循环
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row := 0
	col := 0
	num := 1 // 从1 开始
	for {
		if num == n*n { // 终结条件
			break
		}
		dir := direction[0]
		direction = append(direction[1:], dir)
		for ; num <= n*n; num++ {
			resp[row][col] = num
			row = row + dir[0]
			col = col + dir[1]
			if dir[1] == 1 || dir[0] == 1 {
				if col*dir[1] == n || row*dir[0] == n || resp[row][col] > 0 {
					break
				}
			}
			if dir[1] == -1 || dir[0] == -1 {
				if (col*dir[1] == 0 && row*dir[0] == 0) || resp[row][col] > 0 {
					break
				}
			}
		}
	}
	return resp
}
