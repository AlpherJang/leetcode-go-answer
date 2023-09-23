package main

import "fmt"

// https://leetcode.cn/problems/count-complete-tree-nodes
func main() {
	root := buildTree([]interface{}{1, 2, 3, 4, 5, 6})
	fmt.Println(countNodes(root))
}

func countNodes(root *TreeNode) int {
	queue := []*TreeNode{root}
	count := 0
	for len(queue) > 0 {
		node := queue[0]
		count += 1
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return count
}

func buildTree(data []interface{}) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := &TreeNode{Val: data[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for i < len(data) {
		node := queue[0]
		if node.Left != nil {
			node.Right = &TreeNode{Val: data[i].(int)}
			queue = append(queue, node.Right)
			queue = queue[1:]
		} else {
			node.Left = &TreeNode{Val: data[i].(int)}
			queue = append(queue, node.Left)
		}
		i += 1
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
