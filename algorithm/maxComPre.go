package main

import "fmt"

func main() {
	//s := []string{"","alart","alarm","alammmm","ala"}
	s := []string{"c", "acc", "ccc"}
	//s := []string{}

	//res := maxPre(s)
	res := maxPre1(s)
	fmt.Println(res)
}

func maxPre(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	//res := make([]rune,0)
	needle := []rune(strs[0])
	fmt.Println(len(needle))
	for i := 0; i < len(needle); i++ {
		fmt.Println(i)
		for _, v := range strs {
			fmt.Println(v)
			vl := []rune(v)
			if needle[i] != vl[i] || i > len(vl)-1 {
				return string(needle[0:i])
			}
		}
	}
	return strs[0]
}

func maxPre1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	if strs[0] == "" {
		return ""
	}
	//todo 不需要转成rune slice
	res := []rune(strs[0])
	kk := len(res) - 1
	//fmt.Println(kk)
	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		it := []rune(strs[i])
		for k, v := range it {
			//fmt.Println(k)
			//fmt.Println(v)
			if (k == len(it)-1 || k == kk) && v == res[k] {
				kk = k
				break
			}
			if v != res[k] {
				kk = k - 1
				if kk < 0 {
					return ""
				}
				break
			}
		}
	}
	fmt.Println(kk)
	return ""
	//return string(res[0:kk+1])
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high - low + 1) / 2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
