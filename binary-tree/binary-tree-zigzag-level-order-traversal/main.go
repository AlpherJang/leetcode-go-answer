package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal
func main() {
	fmt.Println(zigzagLevelOrder(buildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	level := 0
	res := make([][]int, 0)
	for ; len(queue) > 0; level++ {
		tmpRes := make([]int, 0)
		tmpArr := make([]*TreeNode, 0)
		for _, item := range queue {
			tmpRes = append(tmpRes, item.Val)
			if item.Left != nil {
				tmpArr = append(tmpArr, item.Left)
			}
			if item.Right != nil {
				tmpArr = append(tmpArr, item.Right)
			}

		}
		queue = tmpArr
		if level%2 == 1 {
			for i, n := 0, len(tmpRes); i < n/2; i++ {
				tmpRes[i], tmpRes[n-1-i] = tmpRes[n-1-i], tmpRes[i]
			}
		}
		res = append(res, tmpRes)
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
