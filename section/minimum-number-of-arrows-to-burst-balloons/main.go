package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons
func main() {
	fmt.Println(findMinArrowShots([][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}))
}

func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	right := points[0][1]
	count := 1
	for _, item := range points {
		if item[0] > right {
			right = item[1]
			count++
		}
	}
	return count
}
