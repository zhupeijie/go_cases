package main

import (
	"reflect"
	"time"
)

var TestC chan int

func main() {
	//dr := common.DemoReflect{}
	TestC = make(chan int, 10)
	go runC()
	time.Sleep(time.Duration(10*10e9))
	tv()
}

func tv()  {
	d := 1
	de := reflect.ValueOf(&d).Elem()
	de.SetInt(11)
	TestC <- d
	de.SetInt(10)

}


func runC()  {
	select {
	case a:= <-TestC:
		a++
	default:

	}
}
