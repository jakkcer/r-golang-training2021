package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var resultNodes []*html.Node
	findElementByTagName := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, q := range name {
				if n.Data == q {
					resultNodes = append(resultNodes, n)
				}
			}
		}
	}
	forEachNode(doc, findElementByTagName, nil)
	return resultNodes
}

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

func main() {
	url := "http://gopl.io"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// 試しに<img>を取得して属性を出力
	images := ElementsByTagName(doc, "img")
	for _, img := range images {
		fmt.Println(img.Data, img.Attr)
	}
}
