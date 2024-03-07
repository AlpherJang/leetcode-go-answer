package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree
func main() {
	printTree(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Print(root.Val)
	fmt.Print(" ")
	printTree(root.Left)
	printTree(root.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Right = buildTree(nums, mid+1, right)
	node.Left = buildTree(nums, left, mid-1)
	return node
}
