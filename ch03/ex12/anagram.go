package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		fmt.Println(isAnagram(os.Args[1], os.Args[2]))
	}
}

func isAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}
