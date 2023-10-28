package main

import "fmt"

// https://leetcode.cn/problems/minimum-genetic-mutation
func main() {
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}

type step struct {
	value string
	step  int
}

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	bankMap := make(map[string]bool, len(bank))
	for _, item := range bank {
		bankMap[item] = true
	}
	if !bankMap[endGene] {
		return -1
	}
	queue := []step{{startGene, 0}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for i, char := range item.value {
			for _, sig := range "ACGT" {
				if sig != char {
					next := item.value[:i] + string(sig) + item.value[i+1:]
					if bankMap[next] {
						if next == endGene {
							return item.step + 1
						}
						delete(bankMap, next)
						queue = append(queue, step{next, item.step + 1})
					}
				}
			}
		}
	}
	return -1
}
