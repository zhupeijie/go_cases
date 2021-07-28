package main

import "fmt"

func main() {
	v1 := []int{1, 3}
	v2 := []int{5, 8}
	l1, l2 := initStr(v1, v2)
	testInsertMid(l1,l2)
	//fmtPt(l1)
	//fmtPt(l2)
	//r := testAppend(l1,l2)
	//fmtPt(r)
	//fmt.Println("------------------")
	//ls := conSli1(l1, l2)
	//fmtPt(ls)
	//test1()
}

func testInsertMid(l1, l2 *ListNode) {
	res := l1
	fmtPt(l1)
	fmtPt(l2)
	fmtPt(res)
	//tmp := res
	l2 = testAppend1(res,l2,false)
	fmtPt(l2)
	//for  {
	//	if tmp.Next==nil {
	//		tmp1 := l2.Next
	//		tmp.Next = l2
	//		tmp.Next.Next =nil
	//		l2 = tmp1
	//		break
	//	}
	//	tmp = tmp.Next
	//}

	fmtPt(res)
	fmtPt(l2)
}

//进行比较操作
func conSli1(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for {
		if l1.Val >= l2.Val {
			res = testAppend(res, l2)
			if l2.Next == nil {
				res = testAppend(res, l1)
				return res
			}

			l2 = l2.Next
			continue
		}
		if l1.Val < l2.Val {
			res = testAppend(res, l1)
			if l1.Next == nil {
				res = testAppend(res, l2)
				return res
			}
			l1 = l1.Next
			continue
		}

	}

}

func fmtPt(v *ListNode) {
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

func designRes(l1 *ListNode, v *ListNode) *ListNode {
	res := l1
	if res == nil {
		return v
	}
	for {
		if res.Next == nil {
			res.Next = v
			return l1
		}
		res = res.Next
	}

}

func initStr(v1 []int, v2 []int) (*ListNode, *ListNode) {
	var l1, l2 ListNode
	l1.Val = v1[0]
	l2.Val = v2[0]
	for k, v := range v1 {
		if k > 0 {
			l1.Next = &ListNode{Val: v}
		}
	}
	for k, v := range v2 {
		if k > 0 {
			l2.Next = &ListNode{Val: v}
		}
	}
	return &l1, &l2
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func conSli(v1 []int, v2 []int) []int {
	if len(v1) == 0 || len(v1) == 0 {
		return append(v1, v2...)
	}
	res := make([]int, 0)
	for _, k := range v1 {
		fmt.Println(k)
		fmt.Println(v2)
		for i := 0; i < len(v2); i++ {
			if k >= v2[i] {
				res = append(res, v2[i])
				continue
			}
			if k < v2[i] {
				res = append(res, k)
				v2 = v2[i:]
				break
			}
		}
	}
	return res
}

func test1() {
	i := 0
	for {
		i++
		fmt.Println("&&", i, "&&")
		if i == 10 {
			return
		}

	}
}

func testAppend(res *ListNode, v *ListNode) *ListNode {
	if res == nil {
		return &ListNode{Val: v.Val}
	}
	//var tmp *ListNode
	tmp := res
	for {
		if tmp.Next == nil {
			tmp.Next = &ListNode{Val: v.Val}
			return res
		}
		tmp = tmp.Next
	}
}

func testAppend1(res *ListNode, v *ListNode,b bool) *ListNode {
	if res == nil {
		res = v
		return v
	}
	//var tmp *ListNode
	tmp := res
	for {
		if tmp.Next == nil {
			if b{
				tmp.Next= v
				return v
			}
			fmt.Println("-----")
			fmtPt(v)
			tmp1 := v
			tmp.Next = tmp1
			fmtPt(tmp1)
			fmtPt(v)
			fmt.Println("-----")
			v = v.Next
			tmp.Next.Next =nil
			fmt.Println(tmp.Next.Next)
			//fmtPt(v)
			fmt.Println("++++")

			return v
		}
		tmp = tmp.Next
	}
}