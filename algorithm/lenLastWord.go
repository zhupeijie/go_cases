package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}
func lengthOfLastWord(s string) int {
	ss := []byte(s)
	num :=0
	for i := len(ss)-1; i >=0 ; i-- {
		fmt.Println("+++++=",string(ss[i]))
		if string(ss[i]) == " "{
			fmt.Println(111)
			return num
		}
		num++
	}
	return num
}
