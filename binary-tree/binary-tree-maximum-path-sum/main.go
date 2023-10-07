package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-maximum-path-sum
func main() {
	fmt.Println(maxPathSum(buildTree([]interface{}{0})))
}

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt32
	reverseTree(root)
	return maxSum
}

func reverseTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(reverseTree(root.Left), 0)
	right := max(reverseTree(root.Right), 0)
	maxSum = max(maxSum, root.Val+left+right)
	return root.Val + max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func buildTree(data []interface{}) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := &TreeNode{Val: data[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(data) {
		node := queue[0]
		queue = queue[1:]
		if data[i] != nil {
			node.Left = &TreeNode{Val: data[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if data[i] != nil {
			node.Right = &TreeNode{Val: data[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
