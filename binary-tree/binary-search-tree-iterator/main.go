package main

import "fmt"


// https://leetcode.cn/problems/binary-search-tree-iterator
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
	data []int
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{data: []int{}}
	iter.reverse(root)
	return iter
}

func (this *BSTIterator) reverse(root *TreeNode) {
	if root == nil {
		return
	}
	this.reverse(root.Left)
	this.data = append(this.data, root.Val)
	this.reverse(root.Right)
}

func (this *BSTIterator) Next() int {
	data := this.data[0]
	this.data = this.data[1:]
	return data
}

func (this *BSTIterator) HasNext() bool {
	return len(this.data) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
