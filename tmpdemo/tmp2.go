package main

import "fmt"

type M1 int32
type M2 int64

func main()  {
	var vM1 M1
	var vM2 M2
	vM1 = 10000
	vM2 = M2(vM1)
	fmt.Println(vM2)
	vM1 = M1(vM2)
	fmt.Println(vM1)
}