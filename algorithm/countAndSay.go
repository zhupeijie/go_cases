package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := cAs1(5)
	fmt.Println(str1)
	// str := "1"
	// for i := 2; i <= 2; i++ {
	// 	str = cAs(str)
	// 	fmt.Println(i, "=---", str)
	// }

}

func cAs(in string) string {
	if in == "" {
		return ""
	}
	ins := []rune(in)
	tmp := make([][]rune, 0)
	for _, v := range ins {
		if len(tmp) == 0 {
			tmp = append(tmp, []rune{v})
			continue
		}
		if v == tmp[len(tmp)-1][0] {
			tmp[len(tmp)-1] = append(tmp[len(tmp)-1], v)
			continue
		}
		tmp = append(tmp, []rune{v})
	}
	ts := ""
	for _, v := range tmp {
		num := len(v)
		ts += strconv.Itoa(num)
		ts += string(v[0])
	}
	return ts
}

func cAs1(step int) string {
	//计数
	pre := "1"
	resS := pre
	for i := 2; i <= step; i++ {
		resS = ""
		for j, start := 0, 0; j <= len(pre); j++ {
			if j == len(pre) || pre[j] != pre[start] {
				resS += strconv.Itoa(j - start)
				resS += string(pre[start])
				start = j
			}
		}
		pre = resS
	}
	return resS
}
