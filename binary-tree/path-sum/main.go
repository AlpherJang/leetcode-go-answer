package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode.cn/problems/path-sum
func main() {
	fmt.Println(hasPathSum(buildTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}), 22))
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

func buildNode(node *TreeNode, data interface{}, queue []*TreeNode) *TreeNode {
	node = &TreeNode{Val: data.(int)}
	queue = append(queue, node)
	return node
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
