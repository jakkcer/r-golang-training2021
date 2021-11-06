package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://pkg.go.dev/golang.org/x/net/html"
	if words, images, err := CountWordsAndImages(url); err == nil {
		fmt.Printf("url: %s\nwords: %d\timages: %d\n", url, words, images)
	} else {
		fmt.Printf("[Error] CountWordsAndImages: %v", err)
	}
}

// CountWordsAndImagesはHTMLドキュメントに対するHTTP GETリクエストをurlへ
// 行い，そのドキュメント内に含まれる単語と画像の数を返します．
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return 0, 0, err
	}
	words, images = countWordsAndImages(doc)
	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// count words
	texts := getTextNodeTexts("", n)
	scanner := bufio.NewScanner(strings.NewReader(texts))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}

	// count images
	images = countImages(0, n)
	return
}

func getTextNodeTexts(texts string, n *html.Node) string {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return texts
		}
	}
	if n.Type == html.TextNode {
		texts += " " + n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = getTextNodeTexts(texts, c)
	}
	return texts
}

func countImages(images int, n *html.Node) int {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		images = countImages(images, c)
	}
	return images
}
