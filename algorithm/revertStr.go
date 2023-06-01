package main

import (
	"fmt"
)

func main() {
	// 输入：s = "PAYPALISHIRING", numRows = 3
	// 输出："PAHNAPLSIIGYIR"
	res := convert1("PAYPALISHIRING", 3)
	fmt.Println(res)
	if res == "PAHNAPLSIIGYIR" {
		fmt.Println(" success ")
	}
}

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	tmp := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		tmp[i] = make([]rune, 0)
	}
	fmt.Println(tmp)
	circle := 2*numRows - 2
	for i, v := range s {
		index := i % circle
		if index < numRows {
			tmp[index] = append(tmp[index], v)
		} else {
			tmp[circle-index] = append(tmp[circle-index], v)
		}
	}
	resStr := ""
	for _, v := range tmp {
		for _, r := range v {
			resStr += string(r)
		}
	}
	return resStr
}
