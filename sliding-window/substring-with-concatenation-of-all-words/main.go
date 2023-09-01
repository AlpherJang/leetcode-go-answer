package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	length, m, n := len(s), len(words), len(words[0])
	totalLength := m * n
	for i := 0; i+totalLength <= length; i++ {
		wordsMap := map[string]int{}
		for _, item := range words {
			wordsMap[item] += 1
		}
		subStr := s[i : i+totalLength]
		match := true
		for j := 0; j < totalLength; j = j + n {
			if wordsMap[subStr[j:j+n]] == 0 {
				match = false
				break
			} else {
				wordsMap[subStr[j:j+n]] -= 1
			}
		}
		if match {
			res = append(res, i)
		}
	}
	return res

}
