package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst
func main() {
	fmt.Println(kthSmallest(buildTree([]interface{}{3, 1, 4, nil, 2}), 1))
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	res, idx := 0, 0
	find := false
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		idx++
		if idx == k {
			res = node.Val
			find = true
		}
		if find {
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
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
