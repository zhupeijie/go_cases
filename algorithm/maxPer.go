package main

import (
	"math"
)

func main() {
}

func findMaxAverage(nums []int, k int) float64 {
	var max float64
	max = math.MinInt32
	for i := 0; i <= len(nums)-k; i++ {
		total := 0
		for j := 0; j < k; j++ {
			total += nums[i+j]
		}
		if float64(total)/float64(k) > max {
			max = float64(total) / float64(k)
		}
	}
	return max
}
