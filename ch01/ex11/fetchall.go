// fetchallはURLを並行に取り出して，時間と大きさをを表示します．
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	urlList := os.Args[1:]
	if len(urlList) == 0 {
		fp, err := os.Open("urlList.txt")
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			urlList = append(urlList, scanner.Text())
		}
	}

	for _, url := range urlList {
		go fetch(url, ch) // ゴルーチンを開始
	}
	for range urlList {
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

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", targetUrl, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, targetUrl)
}
