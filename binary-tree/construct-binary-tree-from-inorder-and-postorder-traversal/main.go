package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	printTree(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	idx := 0
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == root.Val {
			break
		}
	}
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}
