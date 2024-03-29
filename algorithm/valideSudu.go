package main

func isValidSudoku(board [][]byte) bool {
	var row [9][9]int
	var col [9][9]int
	var box [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				if row[i][num] == 1 {
					return false
				} else {
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				if box[i/3*3+j/3][num] == 1 {
					return false
				} else {
					box[i/3*3+j/3][num] = 1
				}
			}
		}
	}
	return true
}
