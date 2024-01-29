package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	sum := 0

	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		sum /= 10
	}

	return resPre.Next
}

func main() {
	l1 := makeListNode([]int{2, 4, 3})
	l2 := makeListNode([]int{5, 6, 4})
	qs := makeListNode([]int{7, 0, 8})
	res := addTwoNumbers(l1, l2)
	for qs != nil {
		if qs.Val != res.Val {
			fmt.Println(false)
			return
		}
		res, qs = res.Next, qs.Next
	}

	fmt.Println(true)
}
