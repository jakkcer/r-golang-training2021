package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input url and id")
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[Error] fetch url: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("[Error] parse html: %v", err)
	}

	resultNode := ElementByID(doc, id)
	if resultNode == nil {
		fmt.Println("No result found")
	} else {
		var attr string
		for _, a := range resultNode.Attr {
			attr += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		fmt.Printf("<%s%s/>\n", resultNode.Data, attr)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var targetElement *html.Node

	getElementByID := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					targetElement = n
					return false
				}
			}
		}
		return true
	}

	forEachNode(doc, getElementByID, nil)
	return targetElement
}

// forEachNodeはnから始まるツリー内の個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します．その二つの関数はオプションです．
// preは子ノードを訪れる前に呼び出され(前順:preorder)，
// postは子ノードを訪れた後に呼び出されます(後順:postorder)．
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if ok := pre(n); !ok {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if ok := post(n); !ok {
			return
		}
	}
}
