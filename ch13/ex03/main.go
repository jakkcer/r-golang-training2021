// Bzipperは入力を読み込み、bzip2の圧縮を行い、結果を書き出します。
package main

import (
	"io"
	"log"
	"os"

	"gobook/ch13/ex03/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
