package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路梳理
1：节点target 距离k的所有子节点
2：距离target 距离k-1父节点的所有子节点
*/

func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	var parentMap map[int]*TreeNode

	parentMap = make(map[int]*TreeNode, 0)

	var findParent func(*TreeNode)

	findParent = func(node *TreeNode) {
		if node.Left != nil {
			parentMap[node.Left.Val] = node
			findParent(node.Left)
		}
		if node.Right != nil {
			parentMap[node.Right.Val] = node
			findParent(node.Right)
		}
	}
	//找到所有节点的父节点
	findParent(root)

	var depSearch func(*TreeNode, *TreeNode, int)

	depSearch = func(node *TreeNode, from *TreeNode, i int) {
		if node == nil {
			return
		}
		if i == k {
			ans = append(ans, node.Val)
			return
		}
		if node.Left != from {
			depSearch(node.Left, node, i+1)
		}
		if node.Right != from {
			depSearch(node.Right, node, i+1)
		}
		if parentMap[node.Val] != from {
			depSearch(parentMap[node.Val], node, i+1)
		}
	}

	depSearch(target, nil, 0)
	return
}
