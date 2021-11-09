package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for k, v := range countClass() {
		fmt.Printf("%s :\t%d\n", k, v)
	}
}

func countClass() map[string]int {
	classNum := make(map[string]int)
	prereqKeys := make([]string, 0, len(prereqs))
	for key := range prereqs {
		prereqKeys = append(prereqKeys, key)
	}

	count := func(class string) []string {
		if v, ok := prereqs[class]; ok {
			for _, c := range v {
				if _, ok := classNum[c]; !ok {
					classNum[c] = 1
				} else {
					classNum[c]++
				}
			}
			return v
		}
		return nil
	}

	breadthFirst(count, prereqKeys)
	return classNum
}

// breadthFirstはworklist内の個々の項目に対してfを呼び出します．
// fから返されたすべての項目はworklistへ追加されます．
// fは，それぞれの項目に対して高々一度しか呼び出されません．
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
