package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("a", "aa"))
}

func minWindow(s string, t string) string {
	sLength, tLength := len(s), len(t)
	if sLength < tLength {
		return ""
	}
	resLen := math.MaxInt32
	resL, resR := -1, -1
	tMap, sMap := map[byte]int{}, map[byte]int{}
	for i := 0; i < tLength; i++ {
		tMap[t[i]]++
	}
	for l, r := 0, 0; r < sLength; r++ {
		if tMap[s[r]] > 0 {
			sMap[s[r]]++
		}
		for l <= r {
			match := true
			for k, v := range tMap {
				if sMap[k] < v {
					match = false
					break
				}
			}
			if !match {
				break
			}
			if r-l+1 < resLen {
				resLen = r - l + 1
				resL, resR = l, r+1
			}
			if _, ok := tMap[s[l]]; ok {
				sMap[s[l]]--
			}
			l++
		}
	}
	if resL == -1 {
		return ""
	}
	return s[resL:resR]
}
