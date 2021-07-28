package main

import "fmt"

func main() {
	res := romanToInt("CDIV")
	//rangS("XIV")
	fmt.Println(res)
}

func romanToInt(s string) int {

	res := 0
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		now := string(s1[i])
		if len(s1) == 1 || i == len(s1)-1 {
			res += getNu(now)
			return res
		}
		next := string(s1[i+1])
		if now == "I" && (next == "V" || next == "X") {
			res = res + getNu(next) - getNu(now)
			i++
			continue
		}
		if now == "X" && (next == "L" || next == "C") {
			res = res + getNu(next) - getNu(now)
			i++
			continue
		}
		if now == "C" && (next == "D" || next == "M") {
			res = res + getNu(next) - getNu(now)
			i++
			continue
		}
		res += getNu(now)
	}

	return res
}

func getNu(n string)int{
	//roman := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	switch n {
	case "M":
		return 1000
	case "D":
		return 500
	case "C":
		return 100
	case "L":
		return 50
	case "X":
		return 10
	case "V":
		return 5
	case "I":
		return 1
	default:
		return 0
	}
}

func rangS(s string) {
	s1 := []rune(s)
	//for _,v:=range s1{
	//	//fmt.Println(string(v))
	//}
	for i := len(s); i > 0; i-- {
		fmt.Println(string(s1[i-1]), "==", i)
	}

}
