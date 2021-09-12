// Echoのテスト
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"hoge"}, "インデックス: 1, 値: hoge\n"},
		{[]string{"hoge", "fuga"}, "インデックス: 1, 値: hoge\nインデックス: 2, 値: fuga\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%q)", test.args)

		out = new(bytes.Buffer)
		if err := echo(test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
