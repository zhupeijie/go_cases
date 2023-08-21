package main

import (
	"fmt"
	"unicode"
)

func main() {
	//result := findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})
	//fmt.Println(result, "==")
	fmt.Println(string('a' + 1))
	//str := shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"})
	//fmt.Println(str)
}

func findShortestSubArray(nums []int) int {
	// 1 先求度
	tmp := make(map[int][]int)

	for i, v := range nums {
		if _, ok := tmp[v]; !ok {
			tmp[v] = []int{1, i, i}
		} else {
			tmp[v][0]++
			tmp[v][2] = i
		}
	}
	m := tmp[nums[0]][0]                          // 最大的度
	minL := tmp[nums[0]][2] - tmp[nums[0]][1] + 1 // 最小的长度子数组
	for v, num := range tmp {
		if v != nums[0] {
			if num[0] > m {
				m = num[0]
				d := num[2] - num[1] + 1
				minL = d
			} else if num[0] == m {
				d := num[2] - num[1] + 1
				if minL > d {
					minL = d
				}
			}

		}
	}

	return minL
}

func shortestCompletingWord(licensePlate string, words []string) string {
	count := make([]int, 26)
	for i := range licensePlate {
		if unicode.IsLetter(rune(licensePlate[i])) {
			count[unicode.ToLower(rune(licensePlate[i]))-'a']++
		}
	}
	str := ""
next:
	for _, word := range words {
		words = words[1:]
		cnt := make([]int, 26)
		for i := range word {
			cnt[word[i]-'a']++
		}
		for i, v := range cnt {
			if v < count[i] {
				goto next
			}
		}
		if str == "" {
			str = word
		}
		if len(word) < len(str) {
			str = word
		}
	}

	return str
}
