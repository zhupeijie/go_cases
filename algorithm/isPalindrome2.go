package main

import "fmt"

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func main() {
	in := []int{1, 1, 2, 1}
	head := initD(in)
	//tmpH := head
	//tmpH1 := head
	//for head != nil {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}
	//rev := reverseList(tmpH)

	//fmt.Println(rev)
	//for rev != nil {
	//	fmt.Println(rev.Val)
	//	rev = rev.Next
	//}
	res := isPalindromeV2(head)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	rev := new(ListNode)
	for head != nil {
		pre = &ListNode{Val: head.Val}
		head = head.Next
		pre.Next = rev.Next
		rev.Next = pre
	}
	return rev.Next
}

func initD(in []int) *ListNode {
	var head *ListNode
	var pre *ListNode
	for _, v := range in {
		if head == nil {
			head = &ListNode{Val: v}
			pre = head
		} else {
			pre.Next = &ListNode{Val: v}
			pre = pre.Next
		}
	}
	return head
}

func isPalindromeV2(head *ListNode) bool {
	if head == nil {
		return false
	}
	head1 := head
	rev := reverseList(head1)
	for head != nil && rev != nil {
		if head1.Val != rev.Val {
			return false
		}
		head1 = head1.Next
		rev = rev.Next
	}
	return true
}
