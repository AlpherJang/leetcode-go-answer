package main

import "fmt"

// https://leetcode.cn/problems/design-add-and-search-words-data-structure
func main() {
	oper := []string{"WordDictionary", "addWord", "addWord", "addWord", "addWord", "search", "search", "addWord", "search", "search", "search", "search", "search", "search"}
	value := [][]string{nil, {"at"}, {"and"}, {"an"}, {"add"}, {"a"}, {".at"}, {"bat"}, {".at"}, {"an."}, {"a.d."}, {"b."}, {"a.d"}, {"."}}
	var handler WordDictionary
	for idx, item := range oper {
		switch item {
		case "WordDictionary":
			handler = Constructor()
			fmt.Print(" null")
		case "addWord":
			handler.AddWord(value[idx][0])
			fmt.Print(" null")
		case "search":
			fmt.Printf(" %v", handler.Search(value[idx][0]))
		}
	}
}

type WordDictionary struct {
	subChar []*WordDictionary
	end     bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		subChar: make([]*WordDictionary, 26),
	}
}

func search(node *WordDictionary, word string) *WordDictionary {
	if len(word) == 0 {
		if node.end {
			return node
		}
		return nil
	}
	char := word[0]
	if char == '.' {
		for _, item := range node.subChar {
			if item != nil {
				res := search(item, word[1:])
				if res != nil && res.end {
					return res
				}
			}
		}
		return nil
	} else {
		char = char - 'a'
		if node.subChar[char] == nil {
			return nil
		}
		return search(node.subChar[char], word[1:])
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, char := range word {
		idx := char - 'a'
		if node.subChar[idx] == nil {
			sub := Constructor()
			node.subChar[idx] = &sub
		}
		node = node.subChar[idx]
	}
	node.end = true
}

func (this *WordDictionary) Search(word string) bool {
	node := search(this, word)
	if node == nil || !node.end {
		return false
	}
	return true
}
