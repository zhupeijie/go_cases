package main

import "fmt"

type myint int32

func main()  {
	var a myint
	var b int32 =100

	a = myint(b)
	fmt.Println(a)
	b = int32(a)
	fmt.Println(b)


}
