package main

import "fmt"

func main() {
	r := general()
	foreachLink(r)
	r1 := reverseList(r)
	foreachLink(r1)
	res := plaind(r,r1)
	fmt.Println(res)
}

func plaind(r1,r2 *ListNode)bool{
	for r1 != nil && r2!= nil{
		fmt.Println("+++++",r1.Val,r2.Val)
		if r1.Val != r2.Val{
			return false
		}
		r1 = r1.Next
		r2 =r2.Next
		fmt.Println("+++++",r1.Val,r2.Val)
	}
	return true
}

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func general() *ListNode {
	tmp := []int{1, 2, 3, 4}
	head := &ListNode{Val: 1,Next: nil}
	root := head
	for _, i := range tmp {
		root.Next = &ListNode{Val: i,Next: nil}
		root = root.Next
	}
	return head.Next
}

func foreachLink(head *ListNode) {
	for head != nil {
		fmt.Println("----", head.Val)
		head = head.Next
	}

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	mid := make([]*ListNode, 0)
	root := head
	for root != nil {
		mid = append(mid, root)
		root = root.Next
	}
	tmp := &ListNode{
		Val: 1,
		Next: nil,
	}
	fmt.Println(mid)

	tmp1 := tmp
	for i := len(mid) - 1; i >= 0; i-- {
		tmp1.Next = mid[i]
		tmp1 = tmp1.Next
		tmp1.Next = nil
	}
	return tmp.Next

}


