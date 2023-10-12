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

//	      236
//	   104    701
//	nil   227 nil  911
//
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst
func main() {
	fmt.Println(getMinimumDifference(buildTree([]interface{}{236, 104, 701, nil, 227, nil, 911})))
}

func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt, -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			res = min(node.Val-pre, res)
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
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
