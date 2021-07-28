package main

import "fmt"

type Meter struct {
	v int64
}

type CentiMeter struct {
	v int64
}

func main()  {
	cm := CentiMeter{
		v: 1000,
	}
	var m Meter
	m = Meter(cm)
	fmt.Println(m)
	cm = CentiMeter(m)
	fmt.Println(cm)
}