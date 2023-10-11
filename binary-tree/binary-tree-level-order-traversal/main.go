package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-level-order-traversal
func main() {
	fmt.Println(levelOrder(buildTree([]interface{}{1})))
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpArr := make([]*TreeNode, 0)
		tmpRes := make([]int, len(queue))
		for idx, node := range queue {
			tmpRes[idx] = node.Val
			if node.Left != nil {
				tmpArr = append(tmpArr, node.Left)
			}
			if node.Right != nil {
				tmpArr = append(tmpArr, node.Right)
			}
		}
		res = append(res, tmpRes)
		queue = tmpArr
	}
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
