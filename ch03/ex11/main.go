package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	b := bytes.Buffer{}
	start := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		start = 1
	}
	end := strings.Index(s, ".")
	if end == -1 {
		end = len(s)
	}
	kasu := s[start:end]
	pre := len(kasu) % 3
	if pre > 0 {
		b.Write([]byte(kasu[:pre]))
		if len(kasu) > pre {
			b.WriteString(",")
		}
	}
	for i, c := range kasu[pre:] {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	b.WriteString(s[end:])
	return b.String()
}
