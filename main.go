package main

import (
	"fmt"
	"reflect"
)

func main() {
	//testStr()
	//testInt()
	testStruct()
	//testSliceStruct()
	//testSet()
	//a := new(MyStruct)
	//a.MyInt = 1
	//a.DemotS.Name = "test"
	//
	//fmt.Println(a)
	//fmt.Println(a.DemotS)
}

type MyInt int

type MyStruct struct {
	MyInt
	DemotS
}

type DemotS struct {
	Name string
}

type DemoReflect struct {
	Name string `mytag:"realname"`
	Age  int    `mytag:"XXX"`
	Addr string `mytag:"hello"`
}

func testStruct() {
	demoS := DemoReflect{Name: "ali", Age: 10}

	demoSType := reflect.TypeOf(demoS)

	dsKind := demoSType.Kind()
	fmt.Println("demoS kind: ", dsKind) //kind:  struct
	dsName := demoSType.Name()
	fmt.Println("demoS name: ", dsName)                      //name:  DemoReflect
	fmt.Println("demoS method num: ", demoSType.NumMethod()) //method:  0

	demoSValue := reflect.ValueOf(demoS)
	fmt.Println("demoS type: ",demoSValue.Type()) //common.DemoReflect

	elem := reflect.ValueOf(&demoS).Elem()
	for i := 0; i < elem.NumField(); i++ {
		item := elem.Field(i)
		fmt.Println(i, " elem kind", item.Kind())
		fmt.Println(i, " elem value", item)
	}
}


func testSliceStruct() {
	demoSliceSt := []DemoReflect{
		{Name: "ali", Age: 10, Addr: "beijing"},
		{Name: "tal", Age: 18, Addr: "beijing"},
	}

	dst := reflect.ValueOf(demoSliceSt)
	fmt.Println("demoSliceSt type is ", dst.Type())
	dstv := dst.Type().Elem()
	fmt.Println("demoSliceSt elem type is ", dstv.Kind())
	for i := 0; i < dstv.NumField(); i++ {
		f := dstv.Field(i)
		fmt.Println(i, " name is: ", f.Name)
		fmt.Println(i, " tag is: ", f.Tag.Get("mytag"))
	}
}

func testSet() {
	var v float32 = 5.21
	pt := reflect.TypeOf(v)
	fmt.Println("kind is:", pt.Kind())

	p := reflect.ValueOf(&v)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	//pe := p.Elem()
	//fmt.Println("settability of pe:", pe.CanSet())
	//pe.SetFloat(13.14)
	//fmt.Println("seted value is:", v)
}

func testInt() {
	var int1 int
	int1 = 65
	iType := reflect.TypeOf(int1)
	fmt.Println(iType.Name(), iType.Align())
	iKind := iType.Kind()
	fmt.Println(iKind)

	iValue := reflect.ValueOf(int1)
	iVType := iValue.Type()
	fmt.Println(iVType)
}

func testStr() {
	//var a string
	//a = "this str"
	//aType := reflect.TypeOf(a)
	//aTypeKind := aType.Kind().String()
	//fmt.Println(aTypeKind)
	//aValue := reflect.ValueOf(a)
	//kind := aValue.Type()
	//fmt.Println(kind)

	type Mystr string
	var ma Mystr
	ma = "my str"
	maValue := reflect.ValueOf(ma)
	mType := maValue.Type()
	mKind := maValue.Kind()

	fmt.Println(mType.String())
	fmt.Println(mType.Name())
	fmt.Println(mKind.String())
}

func testFunc() {
	a := testStr

	aValue := reflect.ValueOf(a)
	fmt.Println(aValue.Kind())
	fmt.Println(aValue.Type())

}

/* flag 十进制：824634187864
低五位表示元数据的类型 flag&31
第6位只读:未导出的未嵌入字段
第7位表示：未导出的嵌入字段，只读
第8位表示：value存储的是指向数据的指针
第9位表示：可以获取到地址
第10位表示：v是个method
10+ 是表示要判断是不是method需要左移的位数.....
只读标记： 0110 0000

0000 0000 0000 0000   0000 0000 1100 0000
0000 0000 0000 0111   0010 0000 0101 1000
*/
