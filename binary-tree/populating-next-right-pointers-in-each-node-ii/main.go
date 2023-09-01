package main

import "fmt"

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	fmt.Println(connect(buildTree([]interface{}{1, 2, 3, 4, 5, nil, 7})))
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == size-1 {
				node.Next = nil
			} else {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

func buildTree(data []interface{}) *Node {
	if len(data) == 0 {
		return nil
	}
	root := &Node{Val: data[0].(int)}
	queue := []*Node{root}
	i := 1
	for len(queue) > 0 && i < len(data) {
		node := queue[0]
		queue = queue[1:]
		if data[i] != nil {
			node.Left = &Node{Val: data[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if data[i] != nil {
			node.Right = &Node{Val: data[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func buildNode(node *Node, data interface{}, queue []*Node) *Node {
	node = &Node{Val: data.(int)}
	queue = append(queue, node)
	return node
}
