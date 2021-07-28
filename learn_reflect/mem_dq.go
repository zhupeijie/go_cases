package main

import (
	"fmt"
	"time"
)

func main()  {

	fmt.Println("start")
	go short()
	go long()
	time.Sleep(1*10e9)
	fmt.Println("end")
}

func  short()  {
	for  {

	}
	fmt.Println("short")
}

func long()  {
	for  {

	}
	fmt.Println("long")
}