// countHTMLNodesはHTMLドキュメントツリー内の要素ごとの個数を数えます
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countHTMLNodes: %v\n", err)
		os.Exit(1)
	}
	nodeCountMap := make(map[string]int, 100)
	countHTMLNodes(nodeCountMap, doc)
	// 出力しておく
	for k, v := range nodeCountMap {
		fmt.Printf("%v: %d個\n", k, v)
	}
}

func countHTMLNodes(nodeMap map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		if v, ok := nodeMap[n.Data]; ok {
			nodeMap[n.Data] = v + 1
		} else {
			nodeMap[n.Data] = 1
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countHTMLNodes(nodeMap, c)
	}
}
