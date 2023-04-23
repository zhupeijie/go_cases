package main

import "fmt"

func main() {
	a := []int{1, 2, 5, 3, 9}
	b := []int{6, 3, 9}

	aL := initData(a)
	bL := initData(b)
	fmtPt1(aL)
	fmtPt1(bL)
	c := sum(aL, bL)
	fmtPt1(c)
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func initData(a []int) *ListNode1 {
	ar := new(ListNode1)
	ar.Val = a[0]
	tmp1 := ar
	for k, v := range a {
		if k > 0 {
			tmp2 := &ListNode1{
				Val: v,
			}
			tmp1.Next = tmp2
			tmp1 = tmp2
		}
	}
	return ar
}

func sum(a *ListNode1, b *ListNode1) *ListNode1 {
	var l1, l2, res ListNode1
	var midtmp int
	for a.Next != nil || b.Next != nil {
		n1, n2 := 0
		if a != nil {

		}
	}

}

func fmtPt1(v *ListNode1) {
	if v == nil {
		fmt.Println("")
		return
	}
	var res []int
	for {
		res = append(res, v.Val)
		if v.Next == nil {
			fmt.Println(">>>>>>>", res)
			return
		}
		v = v.Next
	}
}
