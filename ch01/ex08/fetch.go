// fetchはURLにある内容を表示します．
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	url = addHttpToURL(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return err
	}
	fmt.Fprintf(out, "%s", b)
	return nil
}

func addHttpToURL(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}
