package main
func main () {
	var n int64 = 11
	var b bool = true
	var s string = "test-string-1"
	var a [3]bool = [3]bool{true, false, true}
	var sl []int = []int{1,2,3,4}
	var m map[int]string
	var c chan int
	var in interface{}
	_, _, _, _, _, _, _, _ = n, b, s, a, sl, m, c , in
}
/*

  _type

interface接参数，类型转化的问题，一定是记录下来
interface一定是记录了原有的类型的

断言类型
-------三种类型-----
内置类型

interface 类型

自定义类型
*/

