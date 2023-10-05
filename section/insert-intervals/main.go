package main

import "fmt"

// https://leetcode.cn/problems/insert-interval
func main() {
	fmt.Println(insert([][]int{[]int{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	res := make([][]int, 0)
	inserted := false
	for _, item := range intervals {
		if item[1] < newInterval[0] {
			res = append(res, item)
			continue
		}
		if item[0] > newInterval[1] {
			if !inserted {
				res = append(res, newInterval)
				inserted = true
			}
			res = append(res, item)
			continue
		}
		if item[0] < newInterval[0] {
			newInterval[0] = item[0]
		}
		if item[1] > newInterval[1] {
			newInterval[1] = item[1]
		}
	}
	return res
}
