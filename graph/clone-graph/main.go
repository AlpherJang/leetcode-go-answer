package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

// https://leetcode.cn/problems/clone-graph
func main() {
	printGraph(cloneGraph(buildGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})))
}

func printGraph(node *Node) {
	if node == nil {
		fmt.Println("[]")
	}
	tmpMap := map[int]bool{}
	queue := []*Node{node}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		tmpMap[item.Val] = true
		fmt.Printf("%d :", item.Val)
		for _, tmp := range item.Neighbors {
			fmt.Printf("%d ", tmp.Val)
			if !tmpMap[tmp.Val] {
				queue = append(queue, tmp)
			}
		}
		fmt.Println()
	}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	res := &Node{Val: node.Val}
	tmpMap := map[int]*Node{node.Val: res}
	queue := []*Node{node}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		neighbors := make([]*Node, len(item.Neighbors))
		for idx, sub := range item.Neighbors {
			if tmpMap[sub.Val] != nil {
				neighbors[idx] = tmpMap[sub.Val]
			} else {
				subNode := &Node{Val: sub.Val}
				tmpMap[subNode.Val] = subNode
				neighbors[idx] = subNode
				queue = append(queue, sub)
			}
		}
		tmpMap[item.Val].Neighbors = neighbors
	}
	return res
}

func buildGraph(graphData [][]int) *Node {
	if len(graphData) == 0 {
		return nil
	}
	res := &Node{Val: 1}
	tmpGraph := map[int]*Node{1: res}
	for idx, v := range graphData {
		tmpNeighbors := make([]*Node, len(v))
		for j, item := range v {
			if tmpGraph[item] != nil {
				tmpNeighbors[j] = tmpGraph[item]
			} else {
				sub := &Node{Val: item}
				tmpNeighbors[j] = sub
				tmpGraph[item] = sub
			}
		}
		if tmpGraph[idx+1] != nil {
			tmpGraph[idx+1].Neighbors = tmpNeighbors
		} else {
			tmpGraph[idx+1] = &Node{Val: idx + 1, Neighbors: tmpNeighbors}
		}
	}
	return res
}
