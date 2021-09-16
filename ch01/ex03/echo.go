// 非効率な可能性のあるEchoとstrings.Joinを使ったEcho
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	if err := echoWithForLoop(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echoWithForLoop(args []string) error {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
	return nil
}

func echoWithJoin(args []string) error {
	fmt.Fprintln(out, strings.Join(args, " "))
	return nil
}
