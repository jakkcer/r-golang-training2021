// fetchallはURLを並行に取り出して，時間と大きさをを表示します．
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(targetUrl string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(targetUrl)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	up, err := url.Parse(targetUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %s", err)
		return
	}
	f, err := os.Create(strings.Replace(up.Hostname(), ".", "-", -1) + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return
	}
	defer f.Close()

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", targetUrl, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, targetUrl)
}
