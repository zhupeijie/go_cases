package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	A int
	B string
}

func (t *T) SetA(i int) {
	t.A = i
}

type A struct {
	Name string `json:"name"`
}
type B struct {
	Stu []A `json:"stu"`
}

func main() {
	demoT1()
	//var a A
	////a := A{Name: "zhupejie"}
	//a.Name = "111"
	//a1 := A{Name: "baicai"}
	//
	//b := B{Stu: []A{}}
	//b.Stu = append(b.Stu,a,a1)
	//fmt.Println(b.Stu[0].Name)
	//str,_:= json.Marshal(b)
	//fmt.Println(string(str))
   //"stu":[{"name":"zhupejie"},{"name":"baicai"}]}

	//t := T{23, "skidoo"}
	//s := reflect.ValueOf(&t).Elem()
	//typeOfT := s.Type()
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i,
	//		typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}
	//typePT := reflect.TypeOf(&t)
	//
	//fmt.Printf("%d\n",typePT.NumMethod())
	//for i := 0; i < typePT.NumMethod(); i++ {
	//	m := typePT.Method(i)
	//	fmt.Printf("%d: %s %v\n", m.Index,m.Name,m.Type)
	//}
	//s.Field(0).SetInt(77)
	//s.Field(1).SetString("Sunset Strip")
	//fmt.Println("t is now", t)
	//
	////调用方法/函数
	//m := typePT.Method(0)
	//params := make([]reflect.Value,2)
	//params[0] = reflect.ValueOf(&t)
	//params[1] = reflect.ValueOf(5)
	//m.Func.Call(params)
	//fmt.Println("t is now", t)
}

func demoT1()(b *A)  {
	a  := `{"name":"zhupejie"}`
	err := json.Unmarshal([]byte(a),&b)
	fmt.Println(err)
	return
}

