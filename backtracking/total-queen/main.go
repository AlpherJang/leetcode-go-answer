package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))

}

func totalNQueens(n int) (res int) {
	columns, diagonals1, diagonals2 := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res++
			return
		}
		for col, hasQueen := range columns {
			// d1增加了一个n-1,是为了避免下标溢出
			// 从左上到右下的斜线上所有点的行坐标与列坐标之差是相等的
			// 而统一列上的不同斜线上的点的行坐标与列坐标之差相对于中点对称，绝对值相等，因此在遍历column的计算过程中需要对row及col之差增加一个n-1进行哈希分布到数组中
			d1, d2 := row+n-1-col, col+row
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col], diagonals1[d1], diagonals2[d2] = true, true, true
			backtrack(row + 1)
			columns[col], diagonals1[d1], diagonals2[d2] = false, false, false
		}
	}
	backtrack(0)
	return
}
