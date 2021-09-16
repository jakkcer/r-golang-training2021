// Fetchのテスト
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	path, query, contentType, body string
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
	testWant := "test body"

	descr := fmt.Sprintf("fetch(%q)", testInput)

	out = new(bytes.Buffer)
	if err := fetch(testInput); err != nil {
		t.Errorf("%s failed: %v", descr, err)
	}

	got := out.(*bytes.Buffer).String()
	if got != testWant {
		t.Errorf("%s = %q, want %q", descr, got, testWant)
	}
}

func TestAddHttpToURL(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"http://example.com", "http://example.com"},
		{"hoge.com", "http://hoge.com"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("addHttpToURL(%q)", test.input)
		got := addHttpToURL(test.input)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
