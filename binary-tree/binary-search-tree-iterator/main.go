package main

import "fmt"

func main() {
	op := []string{"BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"}
	data := []interface{}{[]interface{}{7, 3, 15, nil, nil, 9, 20}}
	var handler BSTIterator
	for idx, item := range op {
		switch item {
		case "BSTIterator":
			handler = Constructor(buildTree(data[idx].([]interface{})))
			fmt.Print("null ")
		case "next":
			fmt.Printf("%d ", handler.Next())
		case "hasNext":
			fmt.Printf("%v ", handler.HasNext())
		}
	}
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	data []*TreeNode
	idx  int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		idx:  0,
		data: reverse(root),
	}
}

func reverse(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	left := reverse(root.Left)
	right := reverse(root.Right)
	return append(left, append([]*TreeNode{root}, right...)...)
}

func (this *BSTIterator) Next() int {
	data := this.data[this.idx]
	this.idx++
	return data.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.idx+1 <= len(this.data) {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
