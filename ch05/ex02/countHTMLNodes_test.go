package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCountHTMLNodes(t *testing.T) {
	testWantMap := make(map[string]int, 34)
	testWantMap["head"] = 1
	testWantMap["html"] = 1
	testWantMap["body"] = 1
	testWantMap["main"] = 1
	testWantMap["div"] = 5
	testWantMap["section"] = 2
	testWantMap["p"] = 2

	doc, err := html.Parse(strings.NewReader(testInput))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	gotMap := make(map[string]int, 34)
	countHTMLNodes(gotMap, doc)
	for gotKey, gotValue := range gotMap {
		if v, ok := testWantMap[gotKey]; !ok {
			t.Errorf("%s is not counted", gotKey)
		} else if v != gotValue {
			t.Errorf("%s = %d, want %d", gotKey, gotValue, v)
		}
	}
}

var testInput = `<!DOCTYPE html>
<html lang="en">

<body class="Site">

<main id="page" class="Site-content">
<div class="container">
<div id="nav"></div>
<div class="HomeContainer">
  <section class="HomeSection Hero">
    <p class="Hero-description">
      Linux, macOS, Windows, and more.
    </p>
  </section>

  <section class="HomeSection Playground">
    <div class="Playground-headerContainer">
      <p class="Playground-popout js-playgroundShareEl">Open in Playground</p>
    </div>
    <div class="Playground-inputContainer">
    </div>
  </section>
  
</div>

</div>
</main>
</body>
</html>
`
