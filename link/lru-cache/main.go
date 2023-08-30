package main

import "fmt"

//https://leetcode.cn/problems/lru-cache
func main() {
	var l LRUCache
	opt := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	data := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	for idx, item := range opt {
		switch item {
		case "LRUCache":
			l = Constructor(data[idx][0])
			fmt.Println("null")
		case "put":
			l.Put(data[idx][0], data[idx][1])
			fmt.Println("null")
		case "get":
			data := l.Get(data[idx][0])
			fmt.Println(data)
		}
	}
}

type Node struct {
	Key       int
	Value     int
	Next, Pre *Node
}

type LRUCache struct {
	head     *Node
	last     *Node
	history  map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     &Node{},
		last:     &Node{},
		history:  map[int]*Node{},
		capacity: capacity,
	}
	l.head.Next = l.last
	l.last.Pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.history[key]; ok {
		this.deleteNode(node)
		this.insertHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) insertHead(node *Node) {
	this.head.Next.Pre = node
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next = node
}

func (this *LRUCache) deleteNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.history[key]; ok {
		node.Value = value
		this.deleteNode(node)
		this.insertHead(node)
	} else {
		if this.capacity == len(this.history) {
			rmNode := this.last.Pre
			this.deleteNode(rmNode)
			delete(this.history, rmNode.Key)
		}
		node = &Node{Key: key, Value: value}
		this.insertHead(node)
		this.history[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
