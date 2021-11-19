package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// $fooをf("foo")の結果で置換する関数
func expand(s string, f func(string) string) string {
	replacedWords := make(map[string]string)
	result := s
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if strings.HasPrefix(word, "$") {
			// 二重に置換しないようにチェック
			if _, ok := replacedWords[word]; !ok {
				rw := f(word[1:])
				replacedWords[word] = rw
				result = strings.Replace(result, word, rw, -1)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return result
}

// アルファベットの大文字と小文字を反転する関数
func reverseCase(s string) (result string) {
	for _, c := range s {
		if unicode.IsLower(c) {
			result += strings.ToUpper(string(c))
		} else if unicode.IsUpper(c) {
			result += strings.ToLower(string(c))
		} else {
			result += string(c)
		}
	}
	return result
}
