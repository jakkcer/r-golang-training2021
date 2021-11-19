// findlinksは標準入力から読み込まれたHTMLドキュメント内のリンクを表示します
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visitは，n内で見つかったリンクを一つひとつlinksへ追加し，その結果を返します
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if child := n.FirstChild; child != nil {
		links = visit(links, child)
	}
	if sibling := n.NextSibling; sibling != nil {
		links = visit(links, sibling)
	}
	return links
}
