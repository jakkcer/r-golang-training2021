// Fetchのテスト
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	path, contentType, body string
}

func TestFetch(t *testing.T) {
	response := &Response{
		path:        "/test",
		contentType: "application/json",
		body:        "test body",
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if g, w := r.URL.Path, response.path; g != w {
			t.Errorf("request got path %s, want %s", g, w)
		}
		w.Header().Set("Content-Type", response.contentType)
		io.WriteString(w, response.body)
	}

	// モックサーバ
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	testInput := server.URL + "/test"

	ch := make(chan string)
	go fetch(testInput, ch)
	if _, err := fmt.Fprintln(ioutil.Discard, <-ch); err != nil {
		t.Errorf("fetch failed: %v", err)
	}
}
