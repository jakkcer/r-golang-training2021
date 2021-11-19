package main

import (
	"fmt"
	"os"
)

var prereqs = map[string]map[string]string{
	"algorithms": {"data structures": ""},
	"calculus":   {"linear algebra": ""},
	"compilers": {
		"data structures":       "",
		"formal languages":      "",
		"computer organization": "",
	},
	"data structures":       {"discrete math": ""},
	"databases":             {"data structures": ""},
	"discrete math":         {"intro to programming": ""},
	"formal languages":      {"discrete math": ""},
	"networks":              {"operating systems": ""},
	"operating systems":     {"data structures": "", "computer organization": ""},
	"programming languages": {"data structures": "", "computer organization": ""},
	"linear algebra":        {"calculus": ""},
}

func main() {
	sortedCourses, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] toposort(prereqs) got error:\n\t%v\n", err)
		os.Exit(1)
	}
	for i, course := range sortedCourses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]string) (order []string, err error) {
	seen := make(map[string]bool)
	var visitAll func(items map[string]string, depth int)

	// 循環をチェックするためのスライス．辿ったitemを追加していく
	var loops []string
	visitAll = func(items map[string]string, depth int) {
		for item := range items {
			// depthで呼び出しの深さを数える
			// 最上層ではloopsをリセットする
			if depth == 0 {
				loops = []string{}
			}
			loops = append(loops, item)
			if len(loops) > 1 && item == loops[0] {
				err = fmt.Errorf("loop occurred: %q", loops)
			}
			if !seen[item] {
				seen[item] = true
				visitAll(m[item], depth+1)
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]string, len(m))
	for key := range m {
		keys[key] = ""
	}

	visitAll(keys, 0)
	return order, err
}
