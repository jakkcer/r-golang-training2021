package main

import (
	"fmt"
	"strings"
)

func join(sep string, vals ...string) string {
	if len(vals) == 0 {
		return ""
	}
	result := vals[0]
	for _, val := range vals[1:] {
		result += sep + val
	}
	return result
}

func main() {
	// test := []string{"a", "b", "c"}
	test2 := []string{}
	fmt.Println(strings.Join(test2, ","))

	fmt.Println(join(",", "a"))
}
