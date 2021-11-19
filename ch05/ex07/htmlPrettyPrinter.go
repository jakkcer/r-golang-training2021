package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNodeはnから始まるツリー内の個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します．その二つの関数はオプションです．
// preは子ノードを訪れる前に呼び出され(前順:preorder)，
// postは子ノードを訪れた後に呼び出されます(後順:postorder)．
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attributes string
		for _, a := range n.Attr {
			attributes += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		fmt.Printf("%*s<%s%s", depth*2, "", n.Data, attributes)
		if n.FirstChild == nil {
			fmt.Printf("/")
		}
		fmt.Printf(">\n")
		depth++
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
	if n.Type == html.TextNode {
		if str := strings.TrimSpace(n.Data); str != "" {
			fmt.Printf("%*s%s\n", depth*2, "", str)
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
