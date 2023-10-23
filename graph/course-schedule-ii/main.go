package main

import "fmt"

// https://leetcode.cn/problems/course-schedule-ii
func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	courseDeps := make([]int, numCourses)
	courseEdgeList := make([][]int, numCourses)
	for _, item := range prerequisites {
		courseDeps[item[0]]++
		courseEdgeList[item[1]] = append(courseEdgeList[item[1]], item[0])
	}

	queue := make([]int, 0)
	for idx, item := range courseDeps {
		if item == 0 {
			queue = append(queue, idx)
		}
	}
	var res []int
	for len(queue) > 0 {
		item := queue[0]
		res = append(res, item)
		queue = queue[1:]

		for _, v := range courseEdgeList[item] {
			courseDeps[v]--
			if courseDeps[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(res) != numCourses {
		return []int{}
	}
	return res
}
