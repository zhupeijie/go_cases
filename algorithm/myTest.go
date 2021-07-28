package main

import "fmt"

func main()  {
	a := &Demo{Add: "trsgasg"}
	fmt.Println(a)
	testPtr(a)
	fmt.Println(a)
}

type Demo struct {
	Add string
}

func testPtr(a *Demo)  {
	b := &Demo{Add: "111"}
	*a = *b
	return
}