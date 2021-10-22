package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "printTextNode: %v\n", err)
		os.Exit(1)
	}
	printTextNode(doc)
}

func printTextNode(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}
	if n.Type == html.TextNode {
		if str := strings.TrimSpace(n.Data); str != "" {
			fmt.Fprintf(out, "%s\n", str)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printTextNode(c)
	}
}
