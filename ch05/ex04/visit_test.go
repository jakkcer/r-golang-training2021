package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	testInput := `<!DOCTYPE html>
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
<script src="/lib/godoc/jquery.js" defer></script>

<script src="/lib/godoc/playground.js" defer></script>
<script>var goVersion = "\"go1.16.8\"";</script>
<script src="/lib/godoc/godocs.js" defer></script>
</head>

<body class="Site">

<main id="page" class="Site-content">
<div class="container">
<div class="HomeContainer">
  <a href="/"><img class="Header-logo" src="/lib/godoc/images/go-logo-blue.svg" alt="Go"></a>
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
	testWants := []string{
		"/lib/godoc/style.css",
		"/lib/godoc/jquery.js",
		"/lib/godoc/playground.js",
		"/lib/godoc/godocs.js",
		"/",
		"/lib/godoc/images/go-logo-blue.svg",
	}

	doc, err := html.Parse(strings.NewReader(testInput))
	if err != nil {
		fmt.Fprintf(os.Stderr, "TestVisit: %v\n", err)
		os.Exit(1)
	}
	gotLinks := visit(nil, doc)

	for i, got := range gotLinks {
		if got != testWants[i] {
			t.Errorf("%s does not founded", testWants[i])
		}
	}
}
