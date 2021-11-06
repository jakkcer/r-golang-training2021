package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

var (
	outputDir    string          = "output"
	orgHostNames map[string]bool = make(map[string]bool)
)

// breadthFirstはworklist内の個々の項目に対してfを呼び出します．
// fから返されたすべての項目はworklistへ追加されます．
// fは，それぞれの項目に対して高々一度しか呼び出されません．
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(targetUrl string) []string {
	// 異なるドメインのURLをスキップする
	// NOTE: 問題では保存しないようにとあるが，リンクの抽出もスキップする
	parsedUrl, err := url.Parse(targetUrl)
	if err != nil {
		fmt.Println(err)
	}
	if !orgHostNames[parsedUrl.Hostname()] {
		return nil
	}

	fmt.Println(targetUrl)
	// HTTP GET
	resp, err := http.Get(targetUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// NOTE: 一度読んだら再度セットする必要がある
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "[Error] getting %s: %s", targetUrl, resp.Status)
	}

	// HTMLドキュメントをファイルに保存
	fileName := "index.html"
	if strings.HasSuffix(parsedUrl.Path, ".html") {
		fileName = ""
	}
	saveFilePath := filepath.Join(outputDir, parsedUrl.Path, fileName)
	saveDir := filepath.Dir(saveFilePath)
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, 0777); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if err = savePage(saveFilePath, resp); err != nil {
		fmt.Fprintf(os.Stderr, "[Error] saving %s: %s", targetUrl, err)
	}

	// HTTPレスポンスからリンクを抽出する
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	list, err := extractLinks(resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] extracting links %s: %s", targetUrl, err)
	}
	return list
}

func savePage(fileName string, resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll error: %v", err)
	}
	fp, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("create file error: %v", err)
	}
	defer fp.Close()
	fp.WriteString(string(bodyBytes))
	return nil
}

func extractLinks(resp *http.Response) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[Error] parsing %s as HTML: %v", resp.Request.URL.String(), err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // 不正なURLを無視
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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
	if len(os.Args[1:]) == 0 {
		fmt.Fprintln(os.Stderr, "Please input urls")
		os.Exit(1)
	}
	// 必要に応じて出力用のディレクトリをリセットする
	if _, err := os.Stat(outputDir); !os.IsNotExist(err) {
		if err := os.RemoveAll(outputDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if err := os.Mkdir(outputDir, 0777); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// 大元のドメインを取得しておく
	for _, u := range os.Args[1:] {
		pu, err := url.Parse(u)
		if err != nil {
			fmt.Println(err)
		}
		if _, ok := orgHostNames[pu.Hostname()]; !ok {
			orgHostNames[pu.Hostname()] = true
		}
	}
	// コマンドライン引数から開始して，
	// ウェブを幅優先でクロールする．
	breadthFirst(crawl, os.Args[1:])
}
