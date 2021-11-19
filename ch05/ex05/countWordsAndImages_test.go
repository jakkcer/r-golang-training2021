package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
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
</style>
</head>

<body class="Site">

<main id="page" class="Site-content">
<div class="container">
<div class="HomeContainer">
  <a href="/"><img class="Header-logo" src="/lib/godoc/images/go-logo-blue.svg" alt="Go"></a>
  <img src="test.png"/>
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

  <img src="test.png"/>
  <img src="test.png"/>
  <img src="test.png"/>
  <img src="test.png"/>
  
</div>

</div>
</main>
</body>
</html>
`
	testWants := struct {
		words  int
		images int
	}{15, 6}

	doc, err := html.Parse(strings.NewReader(testInput))
	if err != nil {
		fmt.Fprintf(os.Stderr, "TestCountWordsAndImages: %v\n", err)
		os.Exit(1)
	}
	gotWords, gotImages := countWordsAndImages(doc)

	if gotWords != testWants.words {
		t.Errorf("words shuld be %q, but got %q", testWants.words, gotWords)
	}
	if gotImages != testWants.images {
		t.Errorf("images shuld be %q, but got %q", testWants.images, gotImages)
	}
}
