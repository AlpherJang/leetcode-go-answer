package main

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func flatten(root *TreeNode) {
	list := preTravel(root)
	for i := 1; i < len(list); i++ {
		pre, node := list[i-1], list[i]
		pre.Right = node
		pre.Left = nil
	}
}

func preTravel(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preTravel(root.Left)...)
		list = append(list, preTravel(root.Right)...)
	}
	return list
}
