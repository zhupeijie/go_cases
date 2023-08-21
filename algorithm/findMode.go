package main

func findMode(root *TreeNode) []int {
	base, max, count, resp := 0, 0, 0, []int{}

	update := func(x int) {
		if base == x {
			count++
		} else {
			base = x
			count = 1
		}
		if max == count {
			resp = append(resp, x)
		} else if max < count {
			resp = []int{x}
			max = count
		}
		return
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		update(root.Val)
		dfs(root.Right)
		return
	}
	dfs(root)
	return resp

}

func flowered(flowerbed []int, n int) bool {
	pre := -1
	for i := range flowerbed {
		if n <= 0 {
			return true
		}
		if flowerbed[i] == 1 {
			if pre == -1 {
				n -= i / 2
			} else {
				n -= (i - pre - 2) / 2
			}
			pre = i
		}
	}
	if pre == -1 {
		n -= (len(flowerbed) + 1) / 2
	} else {
		n -= (len(flowerbed) - pre - 1) / 2
	}
	return n <= 0
}
