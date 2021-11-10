package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

// fetchはURLをダウンロードして，ローカルファイルの名前と長さを返します
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() { setNewError(f.Close(), &err) }()

	n, err = io.Copy(f, resp.Body)
	// ファイルを閉じるが，Copyでエラーであればそちらを優先する
	if err != nil {
		return "", 0, err
	}
	return local, n, err
}

func setNewError(newErr error, currentErr *error) {
	if currentErr == nil {
		*currentErr = newErr
	}
}

func main() {
	fetch("http://gopl.io")
}
