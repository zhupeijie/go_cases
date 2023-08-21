package main

func main() {
	in := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(in)
	println(obj.SumRange(0, 2))
}

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			preSum[i] = v
			continue
		}
		preSum[i] = preSum[i-1] + v
	}
	return NumArray{
		data: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.data[right]
	}
	return this.data[right] - this.data[left-1]
}
