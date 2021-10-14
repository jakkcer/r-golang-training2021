package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	input := []byte("あ　　い　　　う")
	output := make([]byte, len(input))
	isMultiSpace := false
	for i := 0; i < len(input); {
		r, size := utf8.DecodeRune(input[i:])
		i += size
		if unicode.IsSpace(r) {
			if isMultiSpace {
				continue
			} else {
				output = append(output, byte(' '))
				isMultiSpace = true
			}
		} else {
			output = append(output, []byte(string(r))...)
			isMultiSpace = false
		}
	}
	fmt.Println(string(output))
}
