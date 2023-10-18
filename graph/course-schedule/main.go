package main

import "fmt"

// https://leetcode.cn/problems/course-schedule
func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	tmpGraph := make([][]int, numCourses)
	visited := make([]int, numCourses)
	res := true

	// 翻转每条边的起始，构建出一个图，图的每条边表示完成这个课程后，后续可继续完成的课程
	for _, item := range prerequisites {
		tmpGraph[item[1]] = append(tmpGraph[item[1]], item[0])
	}

	var dfs func(i int)

	dfs = func(i int) {
		visited[i] = 1
		for _, j := range tmpGraph[i] {
			if visited[j] == 0 {
				dfs(j)
				if !res {
					return
				}
			} else if visited[j] == 1 {
				res = false
				return
			}
		}
		visited[i] = 2
		return
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
		if !res {
			break
		}
	}
	return res
}
