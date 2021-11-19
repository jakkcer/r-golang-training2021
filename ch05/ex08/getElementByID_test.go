package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	targetHTML := `<!DOCTYPE html>
<html lang="en">
  <div id="hoge"/>
  <p id="fuga"/>
  <span id="hoge"/>
</html>`
	tests := []struct {
		inputId     string
		wantElement string
	}{
		{"hoge", "div"},
		{"fuga", "p"},
	}

	doc, err := html.Parse(strings.NewReader(targetHTML))
	if err != nil {
		fmt.Printf("[Error] parse html: %v", err)
	}
	for _, test := range tests {
		descr := fmt.Sprintf("ElementByID(doc, %s)", test.inputId)
		got := ElementByID(doc, test.inputId)
		if got.Data != test.wantElement {
			t.Errorf("%s = %q, but want %q", descr, got.Data, test.wantElement)
		}
	}
}
