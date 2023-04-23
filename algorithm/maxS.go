package main

import "fmt"

func main() {
	s := "babad"
	// s1 := longestPalindrome(s)
	s1 := longest(s)

	//  "#a#c#b#b#c#b#d#s#"
	fmt.Println(s1)
}

//
// func longestPalindrome(s string) string {
// 	start, end := 0, 0
// 	for i := 0; i < len(s); i++ {
// 		l1, r1 := search1(s, i, i)
// 		l2, r2 := search1(s, i, i+1)
// 		if r1-l1 > end-start {
// 			start = l1
// 			end = r1
// 		}
// 		if r2-l2 > end-start {
// 			start = l2
// 			end = r2
// 		}
// 	}
// 	fmt.Println(start, end)
// 	return s[start : end+1]
// }
//
// func search1(s string, left, right int) (int, int) {
// 	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
// 		fmt.Println("=====", left, right)
// 	}
// 	fmt.Println("++++++", left, right)
// 	return left + 1, right - 1
// }

func longest(in string) string {
	start, end := 0, 0
	for i := 0; i < len(in); i++ {
		s1, e1 := sear1(in, i, i)
		s2, e2 := sear1(in, i, i+1)
		if e1-s1 > end-start {
			start = s1
			end = e1
		}
		if e2-s2 > end-start {
			start = s2
			end = e2
		}
	}
	return in[start : end+1]
}

// search longest
func sear1(in string, start, end int) (int, int) {
	for ; start >= 0 && end < len(in) && in[start] == in[end]; start, end = start-1, end+1 {

	}
	return start + 1, end - 1
}
