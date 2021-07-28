package main

import (
	"fmt"
	"reflect"
)

type demo struct {
	Name string
}

func main()  {
	type Book struct {page int}
	x := struct {page int}{123}
	y := Book{123}
	fmt.Println(reflect.DeepEqual(x, y))   // false
	fmt.Println(x == y)                    // true

	z := Book{123}
	fmt.Println(reflect.DeepEqual(&z, &y)) // true
	fmt.Println(&z == &y)             // false



}
