package main

import "fmt"

//https://leetcode.cn/problems/reverse-words-in-a-string
func main() {
	fmt.Println(reverseWords("  hello world  "))
}

func reverseWords(s string) string {
	res := ""
	tmp := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			tmp = string(s[i]) + tmp
		} else if tmp != "" {
			res = getStr(res, tmp)
			tmp = ""
		}
	}
	if tmp != "" {
		res = getStr(res, tmp)
	}
	return res
}

func getStr(res, tmp string) string {
	if len(res) != 0 {
		res += " "
	}
	return res + tmp
}
