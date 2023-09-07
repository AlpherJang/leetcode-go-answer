package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/sum-root-to-leaf-numbers
func main() {
	fmt.Println(sumNumbers(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
}

func sumNumbers(root *TreeNode) int {
	return sum(root, 0)
}

func sum(root *TreeNode, lastNum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return lastNum*10 + root.Val
	}
	levelVal := lastNum*10 + root.Val
	num := sum(root.Left, levelVal) + sum(root.Right, levelVal)
	return num
}
