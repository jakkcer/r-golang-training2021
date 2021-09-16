// fetchはURLにある内容を表示します．
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		if err := fetch(url); err != nil {
			os.Exit(1)
		}
	}
}

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return err
	}

	_, err = io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return err
	}
	return nil
}
