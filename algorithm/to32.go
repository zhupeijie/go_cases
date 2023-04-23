package main

import "fmt"
import "time"

func main() {
	dayNum := time.Now().YearDay()

	ff := to32(dayNum/32)+to32(dayNum%32)
	fmt.Println(dayNum,"---", ff)
	a := to32(29)
	b := to32(164)
	c := to32(26412)
	fmt.Println("---", a, "+++", b, "====", c)
}

var map32 map[int]string

func init() {
	map32 = make(map[int]string)
	map32[1] = "1"
	map32[2] = "2"
	map32[3] = "3"
	map32[4] = "4"
	map32[5] = "5"
	map32[6] = "6"
	map32[7] = "7"
	map32[8] = "8"
	map32[9] = "9"
	map32[10] = "A"
	map32[11] = "B"
	map32[12] = "C"
	map32[13] = "D"
	map32[14] = "E"
	map32[15] = "F"
	map32[16] = "G"
	map32[17] = "H"
	map32[18] = "J"
	map32[19] = "K"
	map32[20] = "L"
	map32[21] = "M"
	map32[22] = "N"
	map32[23] = "P"
	map32[24] = "Q"
	map32[25] = "R"
	map32[26] = "S"
	map32[27] = "T"
	map32[28] = "U"
	map32[29] = "W"
	map32[30] = "X"
	map32[31] = "Y"
	map32[0] = "Z" //Z当0用
}

func to32(source int) string {
	res := ""
	needSlice := make([]int, 0)
	for {
		if source >= 32 {
			needSlice = append(needSlice, source%32)
			source = source / 32
		}
		if source < 32 {
			needSlice = append(needSlice, source)
			break
		}
	}

	for _, v := range needSlice {
		res = map32[v] + res
	}
	return res
}
