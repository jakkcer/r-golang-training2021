// 個々の引数のインデックスと値の組を1行ごとに表示するEcho
package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	if err := echo(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(args []string) error {
	for i, s := range args {
		fmt.Fprintf(out, "インデックス: %v, 値: %v\n", i+1, s)
	}
	return nil
}
