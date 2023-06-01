package main

func main() {

}

func doShudu(board [][]byte) {
	var row [9][9]bool
	var column [9][9]bool
	var box [3][3][9]bool
	var space [][2]int

	// 构造已经完成的数组
	for rowI, rows := range board {
		for columnI, v := range rows {
			if v == '.' {
				space = append(space, [2]int{rowI, columnI})
			} else {
				num := v - '1'
				row[rowI][num] = true
				column[columnI][num] = true
				box[rowI/3][columnI/3][num] = true
			}
		}
	}
	var dsf func(pos int) bool

	dsf = func(pos int) bool {
		if pos == len(space) {
			return true
		}
		i, j := space[pos][0], space[pos][1]
		for num := byte(0); num < 9; num++ {
			if !row[i][num] && !column[j][num] && !box[i/3][j/3][num] {
				row[i][num] = true
				column[j][num] = true
				box[i/3][j/3][num] = true
				board[i][j] = num + '1'
				if dsf(pos + 1) {
					return true
				}
				row[i][num] = false
				column[j][num] = false
				box[i/3][j/3][num] = false
			}
		}
		return false
	}
	dsf(0)
}
