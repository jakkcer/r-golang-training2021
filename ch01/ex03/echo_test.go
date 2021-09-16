// Echoのベンチマーカー
package main

import (
	"bytes"
	"testing"
)

var testInput []string = []string{"hoge", "fuga", "piyo"}

func BenchmarkEchoWithForLoop(b *testing.B) {
	out = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		echoWithForLoop(testInput)
	}
}

func BenchmarkEchoWithJoin(b *testing.B) {
	out = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		echoWithJoin(testInput)
	}
}
