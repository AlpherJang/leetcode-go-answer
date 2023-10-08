package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree
func main() {
	fmt.Println(lowestCommonAncestor(buildTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}), buildNode(5), buildNode(4)))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 查找左子树中是否存在p或q
	leftCheck := lowestCommonAncestor(root.Left, p, q)
	// 查找右子树中是否存在p或q
	rightCheck := lowestCommonAncestor(root.Right, p, q)
	// 如果左子树与右子树均查找到p或q，则p与q的公共祖先节点为当前节点，且当前节点必然是最近的公共祖先节点
	if leftCheck != nil && rightCheck != nil {
		return root
	}
	// 左子树未查找到p或q则两个节点的公共祖先节点为右子树中查找到的节点
	if leftCheck == nil {
		return rightCheck
	}
	return leftCheck
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

func buildNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
