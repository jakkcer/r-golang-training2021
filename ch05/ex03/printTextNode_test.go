package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrintTextNode(t *testing.T) {
	var testInput = `<!DOCTYPE html>
<html lang="en">
<head>
<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">
<title>Golang.com</title>
<script>window.initFuncs = [];</script>
<style>
  body {
    background-color: #00ff00
  }

  .container {
    width: 90%;
    text-align: center;
  }

  .first-sentence {
    font-weight: bold;
  }
</style>
<script>
var _gaq = _gaq || [];
_gaq.push(["_setAccount", "UA-11222381-2"]);
window.trackPageview = function() {
  _gaq.push(["_trackPageview", location.pathname+location.hash]);
};
window.trackPageview();
window.trackEvent = function(category, action, opt_label, opt_value, opt_noninteraction) {
  _gaq.push(["_trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
};
</script>
</head>

<body class="Site">

<main id="page" class="Site-content">
<div class="container">
<div class="HomeContainer">
  <section class="HomeSection Hero">
    <h1>test title h1</h1>
    <p class="Hero-description">
      Linux, macOS, Windows, and more.
    </p>
  </section>

  <section class="HomeSection Playground">
    <h2>test h2</h2>
    <div class="Playground-headerContainer">
      <p class="Playground-popout js-playgroundShareEl">Open in Playground</p>
      <span>spanspanspan</span>
    </div>
  </section>
  
</div>

</div>
</main>
</body>
</html>
`

	var testWant = `Golang.com
test title h1
Linux, macOS, Windows, and more.
test h2
Open in Playground
spanspanspan
`

	doc, err := html.Parse(strings.NewReader(testInput))
	if err != nil {
		fmt.Fprintf(os.Stderr, "printTextNode: %v\n", err)
		os.Exit(1)
	}
	out = new(bytes.Buffer)
	printTextNode(doc)

	got := out.(*bytes.Buffer).String()
	if got != testWant {
		t.Errorf("got:\n%s\nwant:\n%s", got, testWant)
	}
}
