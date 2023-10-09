package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-right-side-view
func main() {
	fmt.Println(averageOfLevels(buildTree([]interface{}{3, 9, 20, 15, 7})))
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	queue := []*TreeNode{root}
	levelLast := 1
	sum := 0
	levelCount := float64(1)
	for len(queue) > 0 {
		item := queue[0]
		sum += item.Val
		levelLast--
		if levelLast == 0 {
			res = append(res, float64(sum)/levelCount)
			levelCount = 0
			sum = 0
		}
		queue = queue[1:]
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
		if levelLast == 0 {
			levelLast = len(queue)
			levelCount = float64(levelLast)
		}
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
