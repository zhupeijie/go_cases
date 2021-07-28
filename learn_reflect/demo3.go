package main

import (
	"fmt"
	"reflect"
)

//用与reflect struct 发射相关
func  main()  {
test()
}

func test()  {
	var a DemoSt
	a.Name="你好"
	a.Age= 11

	tValue := reflect.ValueOf(a)
	tValue.CanAddr()
	fmt.Println(tValue.Kind())
}

type DemoSt struct {
	Name string
	Age int
}

func(d DemoSt) setName(name string)  {
	if name != ""{
		d.Name = name
	}
	return
}

func(d DemoSt) setAge(age int)  {
	if age > 0{
		d.Age = age
	}
}