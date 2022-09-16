package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	n := 3
	s1 := convert(s, n)
	fmt.Println(s1)
}


func convert(s string, numRows int) string {
	resS := make([]string, 0, len(s))
	sByte := []byte(s)
	n := 0
	lo := numRows*2 - 2
	maxY := (len(s)/lo + 1) * (numRows - 1)
	midMa := make([][]string, numRows)
	for i, _ := range midMa {
		midMa[i] = make([]string, maxY)
	}
	for i, b := range sByte {
		n = i / lo
		ix := i - n*lo

		if ix < numRows {
			if i == 7 {
				fmt.Println(string(b), "====")
			}
			midMa[ix][n*(numRows-1)] = string(b)
			continue
		}
		fmt.Println(string(b), i, "----", ix, n, (numRows-1)-(ix-numRows+1), n*(numRows-1)+ix-numRows+1)
		midMa[(numRows-1)-(ix-numRows+1)][n*(numRows-1)+ix-numRows+1] = string(b)
		if i == 7 {
			fmt.Println(string(b), "====")
			fmt.Println("++++++", midMa[1][4])
		}
	}
	fmt.Println("++++++", midMa[1][4])
	a, _ := json.Marshal(midMa)
	fmt.Println(string(a))

	for _, m := range midMa {
		for _, b := range m {
			if b != "" {
				resS = append(resS, b)
			}
		}
	}
	return strings.Join(resS, "")
}
