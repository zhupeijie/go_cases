package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	t := reflect.TypeOf(x)
	fmt.Println(t)
}

