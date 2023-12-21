package main

import "fmt"

// https://leetcode.cn/problems/word-search
func main() {
	fmt.Println(exist([][]byte{{'A'}}, "A"))
}

var directs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if m*n < len(word) {
		return false
	}

	var (
		backTrack func(i, j int, wordIndex int) bool
		find      bool
	)

	backTrack = func(i, j int, wordIndex int) bool {
		if i >= m || j >= n || i < 0 || j < 0 || board[i][j] == ' ' {
			return false
		}
		if board[i][j] != word[wordIndex] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = ' '
		if wordIndex == len(word)-1 {
			find = true
			return true
		}
		for _, direct := range directs {
			if backTrack(i+direct[0], j+direct[1], wordIndex+1) {
				return true
			}
		}
		board[i][j] = tmp
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if find {
				return find
			}
			backTrack(i, j, 0)
		}
	}

	return find
}
