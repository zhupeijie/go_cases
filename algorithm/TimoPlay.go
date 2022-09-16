package main

import "fmt"

func main() {
	in := []int{1,2,3,4,5}
	a := findPoisonedDuration(in,5)
	fmt.Println(a)
}

//次数,中毒时长
func findPoisonedDuration(timeSeries []int, duration int) int {
	var total = 0
	if len(timeSeries) == 0 {
		return total
	}

	for k, v := range timeSeries {
		if k == 0 {
			total += duration
			fmt.Println(total)
			continue
		}
		if v - timeSeries[k-1] >= duration {
			total += duration
		} else {
			total = total + v - timeSeries[k-1]
			fmt.Println(total)
		}
	}
	return total
}
