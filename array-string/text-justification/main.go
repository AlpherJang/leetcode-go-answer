package main

import (
	"fmt"
	"strings"
)

func main() {
	printSlice(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}

func printSlice(data []string) {
	for _, item := range data {
		fmt.Println(item)
	}
}

func fillSpace(n int) string {
	return strings.Repeat(" ", n)
}

func fullJustify(words []string, maxWidth int) []string {
	right, n := 0, len(words)
	res := make([]string, 0)
	for {
		// 每一行的开始都是上一行最后一个字符串的下一个字符串
		left := right
		subLength := 0
		for right < n && subLength+len(words[right])+right-left <= maxWidth {
			subLength += len(words[right])
			right++
		}
		// 最后一行，左对齐
		if right == n {
			s := strings.Join(words[left:], " ")
			res = append(res, s+fillSpace(maxWidth-len(s)))
			return res
		}

		// 当前行包含的字符串个数
		wordsCount := right - left
		// 当前行应包含的空格数
		spacesCount := maxWidth - subLength
		if wordsCount == 1 {
			res = append(res, words[left]+fillSpace(spacesCount))
			continue
		}
		// 计算当前行两个字符串之间平均应有多少空格
		avgSpaces := spacesCount / (wordsCount - 1)
		// 将多余的空格平均到左侧部分字符串之间
		extraSpaces := spacesCount % (wordsCount - 1)
		// 左侧部分字符串之间多一个空格，可以达成条件空格不均匀分配时左侧空格多于右侧
		s1 := strings.Join(words[left:left+extraSpaces+1], fillSpace(avgSpaces+1))
		s2 := strings.Join(words[left+extraSpaces+1:right], fillSpace(avgSpaces))
		res = append(res, s1+fillSpace(avgSpaces)+s2)
	}
}
