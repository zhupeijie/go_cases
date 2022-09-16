package main

import (
	"fmt"
	"math/rand"
)

//给定一个正整数数组 w ，其中 w[i] 代表下标 i 的权重（下标从 0 开始），请写一个函数 pickIndex ，它可以随机地获取下标 i，选取下标
//i 的概率与 w[i] 成正比。
//
//
//
//
// 例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 +
// 3) = 0.75（即，75%）。
//
// 也就是说，选取下标 i 的概率为 w[i] / sum(w) 。
//
//
//
// 示例 1：
//
// 输入：
//["Solution","pickIndex"]
//[[[1]],[]]
//输出：
//[null,0]
//解释：
//Solution solution = new Solution([1]);
//solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。
//
// 示例 2：
//
// 输入：
//["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
//[[[1,3]],[],[],[],[],[]]
//输出：
//[null,1,1,1,1,0]
//解释：
//Solution solution = new Solution([1, 3]);
//solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
//solution.pickIndex(); // 返回 1
//solution.pickIndex(); // 返回 1
//solution.pickIndex(); // 返回 1
//solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。
//
//由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
//[null,1,1,1,1,0]
//[null,1,1,1,1,1]
//[null,1,1,1,0,0]
//[null,1,1,1,0,1]
//[null,1,0,1,0,0]
//......
//诸若此类。
//
//
//
//
// 提示：
//
//
// 1 <= w.length <= 10000
// 1 <= w[i] <= 10^5
// pickIndex 将被调用不超过 10000 次
//
// Related Topics 数学 二分查找 前缀和 随机化
// 👍 141 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func main() {
	source := []int{3, 4, 1, 17}
	obj := Constructor1(source)
	//fmt.Println(obj.SumRight)
	var key0, key1, key2, key3, total float64
	for i := 0; i <= 10010; i++ {
		//	fmt.Println("")
		index := obj.PickIndex()
		//fmt.Println(i, "---index = ", index)
		if index == 0 {
			key0++
		}
		if index == 1 {
			key1++
		}
		if index == 2 {
			key2++
		}
		if index == 3 {
			key3++
		}
		total++
	}
	//fmt.Println("key 0 rate:", key0/total, float64(3)/float64(25))
	//fmt.Println("key 1 rate:", key1/total, float64(4)/float64(25))
	//fmt.Println("key 2 rate:", key2/total, float64(1)/float64(25))
	//fmt.Println("key 3 rate:", key3/total, float64(17)/float64(25))

	randD := make([]int, 0)
	for i := 0; i <= 100; i++ {
		//rand.Seed(time.Now().UnixNano())
		randD = append(randD, rand.Intn(25))
	}
	fmt.Println(randD)
}

type Solution struct {
	SumRight []int
	Sum      int
}

func Constructor1(w []int) Solution {
	r := Solution{Sum: 0, SumRight: []int{}}
	for _, v := range w {
		r.Sum = r.Sum + v
		r.SumRight = append(r.SumRight, r.Sum) //存储每个数据的右 边界
	}
	return r
}

func (this *Solution) PickIndex() int {
	pick := rand.Intn(this.Sum) + 1
	//pick = 8
	//fmt.Println("pick := ", pick)
	//return sort.SearchInts(this.SumRight, pick)
	return findMidSearch(pick, this.SumRight)
}

func findMidSearch(needle int, source []int) int {
	if needle <= source[0] {
		return 0
	}
	if needle > source[0] && needle <= source[1] {
		return 1
	}
	if needle > source[len(source)-2] {
		return len(source) - 1
	}
	middle := len(source)/2 - 1
	if source[middle] >= needle {
		if source[middle-1] < needle {
			return middle
		}
		return findMidSearch(needle, source[0:middle])
	}
	if source[middle] < needle {
		if source[middle+1] >= needle {
			return middle + 1
		}
		return findMidSearch(needle, source[middle:])
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
//leetcode submit region end(Prohibit modification and deletion)
