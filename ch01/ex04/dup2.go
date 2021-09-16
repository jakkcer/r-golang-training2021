// Dup2は入力に2回以上現れた行の数とその行のテキストを表示します．
// 標準入力から読み込むか，名前が指定されたファイルの一覧から読み込みます．
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type DupContent struct {
	files map[string]string
	count int
}

var out io.Writer = os.Stdout

func main() {
	if err := dup2(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func dup2(args []string) error {
	counts := make(map[string]DupContent)
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// 出力を一定にするためにアルファベット順にソートする
	lines := make([]string, 0, len(counts))
	for l := range counts {
		lines = append(lines, l)
	}
	sort.Strings(lines)

	for _, line := range lines {
		content := counts[line]
		if content.count > 1 {
			var files, sep string
			for f := range content.files {
				files += sep + f
				sep = ", "
			}
			fmt.Fprintf(out, "ファイル: %s\t重複数: %d\t内容: %s\n", files, content.count, line)
		}
	}
	return nil
}

func countLines(f *os.File, dupContentMap map[string]DupContent) {
	input := bufio.NewScanner(f)
	initFiles := make(map[string]string)
	initFiles[f.Name()] = ""

	for input.Scan() {
		text := input.Text()
		if content, ok := dupContentMap[text]; !ok {
			dupContentMap[text] = DupContent{
				initFiles,
				1,
			}
		} else {
			newMap := make(map[string]string, len(content.files))
			for k, v := range content.files {
				newMap[k] = v
			}
			newMap[f.Name()] = ""
			dupContentMap[text] = DupContent{
				newMap,
				content.count + 1,
			}
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "NewScanner: %v\n", err)
	}
}
