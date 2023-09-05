package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := buildLink([]int{1, 2})
	printLink(reverseKGroup(list, 2))
}

func buildLink(nums []int) *ListNode {
	head := &ListNode{0, nil}
	cur := head
	for _, v := range nums {
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
	}
	return head.Next
}

func printLink(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	list := []*ListNode{}
	for l := head; l != nil; l = l.Next {
		list = append(list, l)
	}
	for i := 0; i+k <= len(list); i += k {
		for j := 0; j < k/2; j++ {
			list[i+j], list[i+k-j-1] = list[i+k-j-1], list[i+j]
		}
	}
	res := &ListNode{}
	cur := res
	for i := 0; i < len(list); i++ {
		cur.Next = list[i]
		cur = cur.Next
	}
	cur.Next = nil
	return res.Next
}
